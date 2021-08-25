package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sourcegraph/src-cli/internal/exec"
)

type podList struct {
	Items []struct {
		Metadata struct {
			Name string
		}
		Spec struct {
			Containers []struct {
				Name string
			}
		}
	}
}

// Init debug flag on src build
func init() {
	flagSet := flag.NewFlagSet("debug", flag.ExitOnError)

	usageFunc := func() {

		fmt.Fprintf(flag.CommandLine.Output(), `'src debug' gathers and bundles debug data from a Sourcegraph deployment.

USAGE
  src [-v] debug -d=<deployment type> [-out=debug.zip]
`)
	}

	// store value passed to flags
	var (
		//TODO add deployment type selector
		deployment = flagSet.String("d", "", "deployment type")
		outfile    = flagSet.String("out", "debug.zip", "The name of the output zip archive")
	)

	handler := func(args []string) error {
		if err := flagSet.Parse(args); err != nil {
			return err
		}

		//validate out flag
		if *outfile == "" {
			return fmt.Errorf("empty -out flag")
		}
		if strings.HasSuffix(*outfile, ".zip") != true {
			*outfile = *outfile + ".zip"
		}
		// handle deployment flag
		if !((*deployment == "server") || (*deployment == "compose") || (*deployment == "kubernetes")) {
			return fmt.Errorf("must declare -d=<deployment type>, as server, compose, or kubernetes")
		}

		// handle deployment flag
		//switch *deployment {
		//case "server":
		//	fmt.Println("its a single container")
		//case "compose":
		//	fmt.Println("its a docker-compose instance")
		//case "kubernetes":
		//	fmt.Println("its a kubernetes deployment")
		//default:
		//	return fmt.Errorf("must declare -d=<deployment type>, as server, compose, or kubernetes")
		//}

		// open pipe to output file
		out, err := os.OpenFile(*outfile, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0666)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}

		// open zip writer
		defer out.Close()
		zw := zip.NewWriter(out)
		defer zw.Close()

		pods, err := getPods()
		if err != nil {
			return fmt.Errorf("failed to get pods: %w", err)
		}

		// TODO error group, create function type and run as goroutines
		// Runs all kubectl based functions
		if err := archiveEvents(zw); err != nil {
			return fmt.Errorf("running archiveEvents failed: %w", err)
		}
		if err := archiveLogs(zw, pods); err != nil {
			return fmt.Errorf("running archiveLogs failed: %w", err)
		}
		if err := archiveDescribes(zw, pods); err != nil {
			return fmt.Errorf("running archiveDescribes failed: %w", err)
		}
		if err := archiveManifests(zw, pods); err != nil {
			return fmt.Errorf("running archiveManifests failed: %w", err)
		}
		return nil
	}

	// Register the command.
	commands = append(commands, &command{
		aliases:   []string{"debug-dump"},
		flagSet:   flagSet,
		handler:   handler,
		usageFunc: usageFunc,
	})
}

// TODO: make outFile the input from `-out=` flag stored in `outFile`, validate that .zip postpends `outFile`
// TODO: improve logging as kubectl calls run (Desc, Mani)
// TODO: refactor dir structure to be by pod, change logs to .log filetype
// TODO: improve error handling
// TODO: refactor archiveLLogs so that both logs and logs --past are collected in the same loop

func getPods() (podList, error) {
	// Declare buffer type var for kubectl pipe
	var podsBuff bytes.Buffer

	// Get all pod names as json
	getPods := exec.Command("kubectl", "get", "pods", "-l", "deploy=sourcegraph", "-o=json")
	getPods.Stdout = &podsBuff
	getPods.Stderr = os.Stderr
	err := getPods.Run()

	//Declare struct to format decode from podList
	var pods podList

	//Decode json from podList
	if err := json.NewDecoder(&podsBuff).Decode(&pods); err != nil {
		fmt.Errorf("feailed to unmarshall get pods json: %w", err)
	}

	fmt.Println(pods)
	return pods, err
}

func archiveEvents(zw *zip.Writer) error {
	//write events to archive
	k8sEvents, err := zw.Create("outFile/kubectl/events.txt")
	if err != nil {
		return fmt.Errorf("failed to create k8s-events.txt: %w", err)
	}

	//define command to get events
	getEvents := exec.Command("kubectl", "get", "events", "--all-namespaces")
	getEvents.Stdout = k8sEvents
	getEvents.Stderr = os.Stderr

	//get events
	if err := getEvents.Run(); err != nil {
		return fmt.Errorf("running kubectl get events failed: %w", err)
	}

	return nil
}

// gets current pod logs and logs from past containers
func archiveLogs(zw *zip.Writer, pods podList) error {

	// run kubectl logs and write to archive, accounts for containers in pod
	for _, pod := range pods.Items {
		fmt.Println("Archiving logs: ", pod.Metadata.Name, "Containers:", pod.Spec.Containers)
		for _, container := range pod.Spec.Containers {
			logs, err := zw.Create("outFile/kubectl/logs/" + pod.Metadata.Name + "/" + container.Name + ".txt")
			if err != nil {
				return fmt.Errorf("failed to create podLogs.txt: %w", err)
			}

			getLogs := exec.Command("kubectl", "logs", pod.Metadata.Name, "-c", container.Name)
			getLogs.Stdout = logs
			getLogs.Stderr = os.Stderr

			if err := getLogs.Run(); err != nil {
				return fmt.Errorf("running kubectl get logs failed: %w", err)
			}
		}
	}

	// run kubectl logs --previous and write to archive if return not err
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			getPrevLogs := exec.Command("kubectl", "logs", "--previous", pod.Metadata.Name, "-c", container.Name)
			if err := getPrevLogs.Run(); err == nil {
				fmt.Println("Archiving previous logs: ", pod.Metadata.Name, "Containers: ", pod.Spec.Containers)
				prev, err := zw.Create("outFile/kubectl/logs/" + pod.Metadata.Name + "/" + "prev-" + container.Name + ".txt")
				getPrevLogs.Stdout = prev
				if err != nil {
					return fmt.Errorf("failed to create podLogs.txt: %w", err)
				}
			}
		}
	}

	return nil
}

func archiveDescribes(zw *zip.Writer, pods podList) error {
	for _, pod := range pods.Items {
		describes, err := zw.Create("outFile/kubectl/describe/" + pod.Metadata.Name + ".txt")
		if err != nil {
			return fmt.Errorf("failed to create podLogs.txt: %w", err)
		}

		describePod := exec.Command("kubectl", "describe", "pod", pod.Metadata.Name)
		describePod.Stdout = describes
		describePod.Stderr = os.Stderr

		if err := describePod.Run(); err != nil {
			return fmt.Errorf("failer to run describe pod: %w", err)
		}
	}
	return nil
}

func archiveManifests(zw *zip.Writer, pods podList) error {
	for _, pod := range pods.Items {
		manifests, err := zw.Create("outFile/kubectl/manifest/" + pod.Metadata.Name + ".yaml")
		if err != nil {
			return fmt.Errorf("failed to create manifest.yaml: %w", err)
		}

		getManifest := exec.Command("kubectl", "get", "pod", pod.Metadata.Name, "-o", "yaml")
		getManifest.Stdout = manifests
		getManifest.Stderr = os.Stderr

		if err := getManifest.Run(); err != nil {
			fmt.Errorf("failed to get pod yaml: %w", err)
		}
	}
	return nil
}
