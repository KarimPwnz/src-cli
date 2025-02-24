package ui

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dineshappavoo/basex"

	"github.com/sourcegraph/src-cli/internal/batches"
	"github.com/sourcegraph/src-cli/internal/batches/executor"
	"github.com/sourcegraph/src-cli/internal/batches/graphql"
	"github.com/sourcegraph/src-cli/internal/batches/workspace"

	batcheslib "github.com/sourcegraph/sourcegraph/lib/batches"
	"github.com/sourcegraph/sourcegraph/lib/batches/execution"
	"github.com/sourcegraph/sourcegraph/lib/batches/git"
)

var _ ExecUI = &JSONLines{}

type JSONLines struct{}

func (ui *JSONLines) ParsingBatchSpec() {
	logOperationStart(batcheslib.LogEventOperationParsingBatchSpec, &batcheslib.ParsingBatchSpecMetadata{})
}
func (ui *JSONLines) ParsingBatchSpecSuccess() {
	logOperationSuccess(batcheslib.LogEventOperationParsingBatchSpec, &batcheslib.ParsingBatchSpecMetadata{})
}
func (ui *JSONLines) ParsingBatchSpecFailure(err error) {
	logOperationFailure(batcheslib.LogEventOperationParsingBatchSpec, &batcheslib.ParsingBatchSpecMetadata{Error: err.Error()})
}

func (ui *JSONLines) ResolvingNamespace() {
	logOperationStart(batcheslib.LogEventOperationResolvingNamespace, &batcheslib.ResolvingNamespaceMetadata{})
}
func (ui *JSONLines) ResolvingNamespaceSuccess(namespace string) {
	logOperationSuccess(batcheslib.LogEventOperationResolvingNamespace, &batcheslib.ResolvingNamespaceMetadata{NamespaceID: namespace})
}

func (ui *JSONLines) PreparingContainerImages() {
	logOperationStart(batcheslib.LogEventOperationPreparingDockerImages, &batcheslib.PreparingDockerImagesMetadata{})
}
func (ui *JSONLines) PreparingContainerImagesProgress(done, total int) {
	logOperationProgress(batcheslib.LogEventOperationPreparingDockerImages, &batcheslib.PreparingDockerImagesMetadata{Done: done, Total: total})
}
func (ui *JSONLines) PreparingContainerImagesSuccess() {
	logOperationSuccess(batcheslib.LogEventOperationPreparingDockerImages, &batcheslib.PreparingDockerImagesMetadata{})
}

func (ui *JSONLines) DeterminingWorkspaceCreatorType() {
	logOperationStart(batcheslib.LogEventOperationDeterminingWorkspaceType, &batcheslib.DeterminingWorkspaceTypeMetadata{})
}
func (ui *JSONLines) DeterminingWorkspaceCreatorTypeSuccess(wt workspace.CreatorType) {
	var t string
	switch wt {
	case workspace.CreatorTypeVolume:
		t = "VOLUME"
	case workspace.CreatorTypeBind:
		t = "BIND"
	}
	logOperationSuccess(batcheslib.LogEventOperationDeterminingWorkspaceType, &batcheslib.DeterminingWorkspaceTypeMetadata{Type: t})
}

func (ui *JSONLines) ResolvingRepositories() {
	logOperationStart(batcheslib.LogEventOperationResolvingRepositories, &batcheslib.ResolvingRepositoriesMetadata{})
}
func (ui *JSONLines) ResolvingRepositoriesDone(repos []*graphql.Repository, unsupported batches.UnsupportedRepoSet, ignored batches.IgnoredRepoSet) {
	logOperationSuccess(batcheslib.LogEventOperationResolvingRepositories, &batcheslib.ResolvingRepositoriesMetadata{
		Unsupported: len(unsupported),
		Ignored:     len(ignored),
		Count:       len(repos),
	})
}

func (ui *JSONLines) DeterminingWorkspaces() {
	logOperationStart(batcheslib.LogEventOperationDeterminingWorkspaces, &batcheslib.DeterminingWorkspacesMetadata{})
}
func (ui *JSONLines) DeterminingWorkspacesSuccess(num int) {
	logOperationSuccess(batcheslib.LogEventOperationDeterminingWorkspaces, &batcheslib.DeterminingWorkspacesMetadata{Count: num})
}

func (ui *JSONLines) CheckingCache() {
	logOperationStart(batcheslib.LogEventOperationCheckingCache, &batcheslib.CheckingCacheMetadata{})
}
func (ui *JSONLines) CheckingCacheSuccess(cachedSpecsFound int, tasksToExecute int) {
	logOperationSuccess(batcheslib.LogEventOperationCheckingCache, &batcheslib.CheckingCacheMetadata{
		CachedSpecsFound: cachedSpecsFound,
		TasksToExecute:   tasksToExecute,
	})
}

func (ui *JSONLines) ExecutingTasks(verbose bool, parallelism int) executor.TaskExecutionUI {
	return &taskExecutionJSONLines{verbose: verbose, parallelism: parallelism}
}

func (ui *JSONLines) ExecutingTasksSkippingErrors(err error) {
	logOperationSuccess(batcheslib.LogEventOperationExecutingTasks, &batcheslib.ExecutingTasksMetadata{
		Skipped: true,
		Error:   err.Error(),
	})
}

func (ui *JSONLines) LogFilesKept(files []string) {
	for _, path := range files {
		logOperationSuccess(batcheslib.LogEventOperationLogFileKept, &batcheslib.LogFileKeptMetadata{Path: path})
	}
}

func (ui *JSONLines) NoChangesetSpecs() {
	ui.UploadingChangesetSpecsSuccess([]graphql.ChangesetSpecID{})
}

func (ui *JSONLines) UploadingChangesetSpecs(num int) {
	logOperationStart(batcheslib.LogEventOperationUploadingChangesetSpecs, &batcheslib.UploadingChangesetSpecsMetadata{
		Done:  0,
		Total: num,
	})
}

func (ui *JSONLines) UploadingChangesetSpecsProgress(done, total int) {
	logOperationProgress(batcheslib.LogEventOperationUploadingChangesetSpecs, &batcheslib.UploadingChangesetSpecsMetadata{
		Done:  done,
		Total: total,
	})
}

func (ui *JSONLines) UploadingChangesetSpecsSuccess(ids []graphql.ChangesetSpecID) {
	sIDs := make([]string, len(ids))
	for i, id := range ids {
		sIDs[i] = string(id)
	}
	logOperationSuccess(batcheslib.LogEventOperationUploadingChangesetSpecs, &batcheslib.UploadingChangesetSpecsMetadata{
		Done:  len(ids),
		Total: len(ids),
		IDs:   sIDs,
	})
}

func (ui *JSONLines) CreatingBatchSpec() {
	logOperationStart(batcheslib.LogEventOperationCreatingBatchSpec, &batcheslib.CreatingBatchSpecMetadata{})
}

func (ui *JSONLines) CreatingBatchSpecSuccess(batchSpecURL string) {
	logOperationSuccess(batcheslib.LogEventOperationCreatingBatchSpec, &batcheslib.CreatingBatchSpecMetadata{
		PreviewURL: batchSpecURL,
	})
}

func (ui *JSONLines) CreatingBatchSpecError(err error) error {
	logOperationFailure(batcheslib.LogEventOperationCreatingBatchSpec, &batcheslib.CreatingBatchSpecMetadata{})
	return err
}

func (ui *JSONLines) PreviewBatchSpec(batchSpecURL string) {
	// Covered by CreatingBatchSpecSuccess.
}

func (ui *JSONLines) ApplyingBatchSpec() {
	logOperationStart(batcheslib.LogEventOperationApplyingBatchSpec, &batcheslib.ApplyingBatchSpecMetadata{})
}

func (ui *JSONLines) ApplyingBatchSpecSuccess(batchChangeURL string) {
	logOperationSuccess(batcheslib.LogEventOperationApplyingBatchSpec, &batcheslib.ApplyingBatchSpecMetadata{BatchChangeURL: batchChangeURL})
}

func (ui *JSONLines) ExecutionError(err error) {
	logOperationFailure(batcheslib.LogEventOperationBatchSpecExecution, &batcheslib.BatchSpecExecutionMetadata{Error: err.Error()})
}

var _ executor.JSONCacheWriter = &JSONLines{}

func (ui *JSONLines) WriteExecutionResult(key string, value execution.Result) {
	logOperationSuccess(batcheslib.LogEventOperationCacheResult, &batcheslib.CacheResultMetadata{
		Key:   key,
		Value: value,
	})
}

func (ui *JSONLines) WriteAfterStepResult(key string, value execution.AfterStepResult) {
	logOperationSuccess(batcheslib.LogEventOperationCacheAfterStepResult, &batcheslib.CacheAfterStepResultMetadata{
		Key:   key,
		Value: value,
	})
}

type taskExecutionJSONLines struct {
	verbose     bool
	parallelism int

	linesTasks map[*executor.Task]batcheslib.JSONLinesTask
}

// seededRand is used in randomID() to generate a "random" number.
var seededRand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

// randomID generates a random ID to be used for identifiers in tasks.
func randomID() (string, error) {
	return basex.Encode(strconv.Itoa(seededRand.Int()))
}

func (ui *taskExecutionJSONLines) Start(tasks []*executor.Task) {
	ui.linesTasks = make(map[*executor.Task]batcheslib.JSONLinesTask, len(tasks))
	linesTasks := []batcheslib.JSONLinesTask{}
	for _, t := range tasks {
		id, err := randomID()
		if err != nil {
			panic(err)
		}
		linesTask := batcheslib.JSONLinesTask{
			ID:                     id,
			Repository:             t.Repository.Name,
			Workspace:              t.Path,
			Steps:                  t.Steps,
			CachedStepResultsFound: t.CachedResultFound,
			StartStep:              t.CachedResult.StepIndex,
		}
		ui.linesTasks[t] = linesTask
		linesTasks = append(linesTasks, linesTask)
	}

	logOperationStart(batcheslib.LogEventOperationExecutingTasks, &batcheslib.ExecutingTasksMetadata{
		Tasks: linesTasks,
	})
}
func (ui *taskExecutionJSONLines) Success() {
	logOperationSuccess(batcheslib.LogEventOperationExecutingTasks, &batcheslib.ExecutingTasksMetadata{})
}

func (ui *taskExecutionJSONLines) Failed(err error) {
	logOperationFailure(batcheslib.LogEventOperationExecutingTasks, &batcheslib.ExecutingTasksMetadata{Error: err.Error()})
}

func (ui *taskExecutionJSONLines) TaskStarted(task *executor.Task) {
	lt, ok := ui.linesTasks[task]
	if !ok {
		panic("unknown task started")
	}

	logOperationStart(batcheslib.LogEventOperationExecutingTask, &batcheslib.ExecutingTaskMetadata{TaskID: lt.ID})
}

func (ui *taskExecutionJSONLines) TaskFinished(task *executor.Task, err error) {
	lt, ok := ui.linesTasks[task]
	if !ok {
		panic("unknown task started")
	}

	if err != nil {
		logOperationFailure(batcheslib.LogEventOperationExecutingTask, &batcheslib.ExecutingTaskMetadata{
			TaskID: lt.ID,
			Error:  err.Error(),
		})
		return
	}

	logOperationSuccess(batcheslib.LogEventOperationExecutingTask, &batcheslib.ExecutingTaskMetadata{TaskID: lt.ID})
}

func (ui *taskExecutionJSONLines) TaskChangesetSpecsBuilt(task *executor.Task, specs []*batcheslib.ChangesetSpec) {
	lt, ok := ui.linesTasks[task]
	if !ok {
		panic("unknown task started")
	}

	logOperationSuccess(batcheslib.LogEventOperationTaskBuildChangesetSpecs, &batcheslib.TaskBuildChangesetSpecsMetadata{TaskID: lt.ID})
}

func (ui *taskExecutionJSONLines) StepsExecutionUI(task *executor.Task) executor.StepsExecutionUI {
	lt, ok := ui.linesTasks[task]
	if !ok {
		panic("unknown task started")
	}

	return &stepsExecutionJSONLines{linesTask: &lt}
}

type stepsExecutionJSONLines struct {
	linesTask *batcheslib.JSONLinesTask
}

const stepFlushDuration = 500 * time.Millisecond

func (ui *stepsExecutionJSONLines) ArchiveDownloadStarted() {
	logOperationStart(batcheslib.LogEventOperationTaskDownloadingArchive, &batcheslib.TaskDownloadingArchiveMetadata{TaskID: ui.linesTask.ID})
}
func (ui *stepsExecutionJSONLines) ArchiveDownloadFinished(err error) {
	if err != nil {
		logOperationFailure(batcheslib.LogEventOperationTaskDownloadingArchive, &batcheslib.TaskDownloadingArchiveMetadata{TaskID: ui.linesTask.ID, Error: err.Error()})
	} else {
		logOperationSuccess(batcheslib.LogEventOperationTaskDownloadingArchive, &batcheslib.TaskDownloadingArchiveMetadata{TaskID: ui.linesTask.ID})
	}
}

func (ui *stepsExecutionJSONLines) WorkspaceInitializationStarted() {
	logOperationStart(batcheslib.LogEventOperationTaskInitializingWorkspace, &batcheslib.TaskInitializingWorkspaceMetadata{TaskID: ui.linesTask.ID})
}
func (ui *stepsExecutionJSONLines) WorkspaceInitializationFinished() {
	logOperationSuccess(batcheslib.LogEventOperationTaskInitializingWorkspace, &batcheslib.TaskInitializingWorkspaceMetadata{TaskID: ui.linesTask.ID})
}

func (ui *stepsExecutionJSONLines) SkippingStepsUpto(startStep int) {
	logOperationProgress(batcheslib.LogEventOperationTaskSkippingSteps, &batcheslib.TaskSkippingStepsMetadata{TaskID: ui.linesTask.ID, StartStep: startStep})
}

func (ui *stepsExecutionJSONLines) StepSkipped(step int) {
	logOperationProgress(batcheslib.LogEventOperationTaskStepSkipped, &batcheslib.TaskStepSkippedMetadata{TaskID: ui.linesTask.ID, Step: step})
}

func (ui *stepsExecutionJSONLines) StepPreparingStart(step int) {
	logOperationStart(batcheslib.LogEventOperationTaskPreparingStep, &batcheslib.TaskPreparingStepMetadata{TaskID: ui.linesTask.ID, Step: step})
}
func (ui *stepsExecutionJSONLines) StepPreparingSuccess(step int) {
	logOperationSuccess(batcheslib.LogEventOperationTaskPreparingStep, &batcheslib.TaskPreparingStepMetadata{TaskID: ui.linesTask.ID, Step: step})
}
func (ui *stepsExecutionJSONLines) StepPreparingFailed(step int, err error) {
	logOperationFailure(batcheslib.LogEventOperationTaskPreparingStep, &batcheslib.TaskPreparingStepMetadata{TaskID: ui.linesTask.ID, Step: step, Error: err.Error()})
}

func (ui *stepsExecutionJSONLines) StepStarted(step int, runScript string, env map[string]string) {
	logOperationStart(batcheslib.LogEventOperationTaskStep, &batcheslib.TaskStepMetadata{TaskID: ui.linesTask.ID, Step: step, Env: env})
}

func (ui *stepsExecutionJSONLines) StepOutputWriter(ctx context.Context, task *executor.Task, step int) executor.StepOutputWriter {
	sink := func(data string) {
		logOperationProgress(
			batcheslib.LogEventOperationTaskStep,
			&batcheslib.TaskStepMetadata{
				TaskID: ui.linesTask.ID,
				Step:   step,
				Out:    data,
			},
		)
	}
	return NewIntervalProcessWriter(ctx, stepFlushDuration, sink)
}

func (ui *stepsExecutionJSONLines) StepFinished(step int, diff string, changes *git.Changes, outputs map[string]interface{}) {
	logOperationSuccess(
		batcheslib.LogEventOperationTaskStep,
		&batcheslib.TaskStepMetadata{
			TaskID:  ui.linesTask.ID,
			Step:    step,
			Diff:    diff,
			Outputs: outputs,
		},
	)
}

func (ui *stepsExecutionJSONLines) StepFailed(step int, err error, exitCode int) {
	logOperationFailure(
		batcheslib.LogEventOperationTaskStep,
		&batcheslib.TaskStepMetadata{
			TaskID:   ui.linesTask.ID,
			Step:     step,
			Error:    err.Error(),
			ExitCode: exitCode,
		},
	)
}

func (ui *stepsExecutionJSONLines) CalculatingDiffStarted() {
	logOperationStart(batcheslib.LogEventOperationTaskCalculatingDiff, &batcheslib.TaskCalculatingDiffMetadata{TaskID: ui.linesTask.ID})
}
func (ui *stepsExecutionJSONLines) CalculatingDiffFinished() {
	logOperationSuccess(batcheslib.LogEventOperationTaskCalculatingDiff, &batcheslib.TaskCalculatingDiffMetadata{TaskID: ui.linesTask.ID})
}

func logOperationStart(op batcheslib.LogEventOperation, metadata interface{}) {
	logEvent(batcheslib.LogEvent{Operation: op, Status: batcheslib.LogEventStatusStarted, Metadata: metadata})
}

func logOperationSuccess(op batcheslib.LogEventOperation, metadata interface{}) {
	logEvent(batcheslib.LogEvent{Operation: op, Status: batcheslib.LogEventStatusSuccess, Metadata: metadata})
}

func logOperationFailure(op batcheslib.LogEventOperation, metadata interface{}) {
	logEvent(batcheslib.LogEvent{Operation: op, Status: batcheslib.LogEventStatusFailure, Metadata: metadata})
}

func logOperationProgress(op batcheslib.LogEventOperation, metadata interface{}) {
	logEvent(batcheslib.LogEvent{Operation: op, Status: batcheslib.LogEventStatusProgress, Metadata: metadata})
}

func logEvent(e batcheslib.LogEvent) {
	e.Timestamp = time.Now().UTC().Truncate(time.Millisecond)
	err := json.NewEncoder(os.Stdout).Encode(e)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
