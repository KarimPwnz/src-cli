<!--
###################################### READ ME ###########################################
### This changelog should always be read on `main` branch. Its contents on version     ###
### branches do not necessarily reflect the changes that have gone into that branch.   ###
##########################################################################################
-->

# Changelog

All notable changes to `src-cli` are documented in this file.

## Unreleased

### Added

### Changed

### Fixed

### Removed

## 3.35.2

### Fixed

- `src batch validate` and `src batch repos` now accept `-allow-unsupported` and `-force-override-ignore` and don't fail on specs using `files` anymore.

## 3.35.1

### Changed

- `src validate` has been updated to work with Sourcegraph 3.35's changed CSRF behaviour. [#673](https://github.com/sourcegraph/src-cli/pull/673)

## 3.35.0

### Added

- Individual batch changes can publish multiple changesets to the same repository by specifying multiple target branches using the [`on.branches`](https://docs.sourcegraph.com/batch_changes/references/batch_spec_yaml_reference#on-repository) attribute. [#25228](https://github.com/sourcegraph/sourcegraph/issues/25228)

### Changed

- `src batch [preview|apply]` will now use the faster volume workspace by default for macOS systems using arm64 processors (aka "Apple Silicon")

### Fixed

- Excess newlines in between outputs in logfiles written when `--keep-logs` is used have been fixed. [#665](https://github.com/sourcegraph/src-cli/pull/665)
- `src` would sometimes panic when Ctrl-C was pressed while executing batch change steps due to a bug in the library used to render the execution progress bars. This has been fixed. [🤘 #666](https://github.com/sourcegraph/src-cli/pull/666)
- In batch changes, when using `workspaces` src would incorrectly treat `in` not being set as _don't match anything_. This is fixed and it matches everything like `*`. [#664](https://github.com/sourcegraph/src-cli/pull/664)

## 3.34.1

### Added

- For internal use only: Allow exec mode to read step cache files from the current working directory.

### Fixed

- For internal use only: Fix an off-by-one error in the JSON log UI.

## 3.34.0

### Added

- Bare repositories can now be served with `src serve-git`. Thanks to [Andreas Rammhold](https://github.com/andir) for the contribution! [#639](https://github.com/sourcegraph/src-cli/pull/639)

## 3.33.8

### Fixed

- For internal use only: Switch to single workspaces only in payload.

## 3.33.7

### Added

- A `-client-only` flag has been added to `src version`. When set, only the local src-cli's version will be printed and no attempt is made to fetch the recommended version from the configured Sourcegraph server.

## 3.33.6

### Added

- Requests to Sourcegraph will now include the operating system and architecture `src` is running on by default. To disable this, set the `SRC_DISABLE_USER_AGENT_TELEMETRY` environment variable to any non-empty string, or provide the `-user-agent-telemetry=false` flag on the command line. [#15769](https://github.com/sourcegraph/sourcegraph/issues/15769)

## 3.33.5

### Fixed

- `src search` will not fail anymore on systems where the pager `less` is not installed. [#644](https://github.com/sourcegraph/src-cli/pull/644)

## 3.33.4

- For internal use only: always log uploading changesets event.

## 3.33.3

### Fixed

- Removed debug output that could lead to glitches in the `src batch [preview|apply]` TUI.

## 3.33.2

### Changed

- For internal use only: `src batch exec` does not evaluate `importChangesets` in batch specs anymore, since that's already done server-side.

## 3.33.1

### Fixed

- Fixes a nil-panic that could be caused when `src batch [preview|apply]` would encounter a repository that was currently being cloned or is empty.

## 3.33.0

### Fixed

- Importing changesets won't fail validation when uploading to Sourcegraph 3.33.

## 3.32.1

### Fixed

- Fixes an issue where src-cli would panic when importing existing changesets.

## 3.32.0

### Added

- For internal use only: the `src batch exec` command executes the provided batch spec in the given workspaces.

### Changed

- For internal use only: when `src batch [preview|apply|exec]` are executed in `-text-only` mode, command output on stdout/stderr will be logged in the same message, with each line prefixed accordingly. [#619](https://github.com/sourcegraph/src-cli/pull/619)

### Fixed

- `src batch repos` failed with a template error in src-cli 3.31.1 and 3.32.0. This has been fixed. [#625](https://github.com/sourcegraph/src-cli/pull/625)
- The `src campaign` and `src campaigns` commands have been removed in favor of `src batch`. [#626](https://github.com/sourcegraph/src-cli/pull/626)

## 3.31.1

### Changed

- For internal use only: JSON log UI has been extended.

### Fixed

- `src search -stream` displayed the number of lines that contain matches instead of the number of matches.
- For internal use only: the `EXECUTING_TASKS` JSON log line now always contains an array of `tasks` instead of possibly having `null` as the `tasks` value.
- src-cli is now built using Go version 1.17.

## 3.31.0

### Changed

- `src batch new` now omits the publish flag, if the Sourcegraph instance supports publish from GUI.

## 3.30.5

### Added

- `src batch validate` now accepts the standard flags used to control communication with Sourcegraph: `-dump-requests`, `-get-curl`, `-insecure-skip-verify`, and `-trace`. [#577](https://github.com/sourcegraph/src-cli/pull/577)

### Fixed

- `src batch validate` would fail to validate batch specs that use features that depend on specific versions of Sourcegraph, such as workspaces. This has been fixed. [#576](https://github.com/sourcegraph/src-cli/issues/576)

## 3.30.4

### Added

- Releases are now built for `arm64` architectures too.

## 3.30.3

### Fixed

- For internal use only: `-text-only` silently ignored an error when trying to print log messages and did not print a `EXECUTING_TASKS` message.

## 3.30.2

### Changed

- For internal use only: `-text-only` now contains detailed information about task execution. [#571](https://github.com/sourcegraph/src-cli/pull/571)
- As part of the above: the TUI of `src batch [preview|apply]` has been reworked and should now feel snappier.

## 3.30.1

### Fixed

- The per-step caching of batch spec execution results was broken when re-execution could use the cached results of a step and that step was the only one left to execute. That resulted in empty diffs being uploaded. This is now fixed. [#567](https://github.com/sourcegraph/src-cli/pull/567)

## 3.30.0

Re-release of 3.29.3 for Sourcegraph 3.30.

## 3.29.3

### Fixed

- `count:all` was not supported in repository search queries for batch changes. This is now fixed. [#566](https://github.com/sourcegraph/src-cli/pull/566)
- For internal use only: `-text-only` received a few tiny fixes for pluralisation and missing log messages. [#565](https://github.com/sourcegraph/src-cli/pull/565)

## 3.29.2

### Added

- Starting with Sourcegraph 3.30.0, the `published` field is optional in batch specs. If omitted, the publication state will be controlled through the Batch Changes UI. [#538](https://github.com/sourcegraph/src-cli/pull/538)
- For internal use only: `-text-only` flag added to `src batch [apply|preview]`. [#562](https://github.com/sourcegraph/src-cli/pull/562)

## 3.29.1

### Added

- LSIF uploads now respect the `-insecure-skip-verify` flag to insecurely (surprise!) skip TLS certificate validation when communicating with Sourcegraph. [#559](https://github.com/sourcegraph/src-cli/pull/559)

### Fixed

- Various terminal handling fixes, especially for Batch Changes users with regards to resizing terminals and Windows support. [#560](https://github.com/sourcegraph/src-cli/pull/560)

## 3.29.0

### Removed

- The `src repos enable|disable` commands were removed as they are no longer supported.

## 3.28.3

### Fixed

- Cached step results produced by `src batch [apply|preview]` are now properly cleared when using the `-clear-cache` command line flag.

## 3.28.2

### Fixed

- The step-wise caching for `src batch [apply|preview]` introduced in 3.28.1 could break if a cached diff contained quoted. This fixes the application by disabling any unquoting/expansion.
- A regression was introduced in 3.28.1 that broke the UI for `src batch [apply|preview]` and lead to the execution of steps looking like it got stuck in the first repository.

## 3.28.1

### Changed

- `src batch [apply|preview]` now cache the results of each step when executing a batch spec. That can make re-execution a lot faster when only a subset of the steps has been changed. [#540](https://github.com/sourcegraph/src-cli/pull/540)

### Fixed

- `src serve-git` can now handle gzip-encoded requests. [#547](https://github.com/sourcegraph/src-cli/pull/547)

## 3.28.0

- This release is identical to 3.27.1, and is simply a version number bump to correspond with the forthcoming release of Sourcegraph 3.28.

## 3.27.1

### Added

- Starting with Sourcegraph 3.28.0 batch spec `steps` can contain an `if: <template string>` attribute that determines whether the given step will be executed. [#520](https://github.com/sourcegraph/src-cli/pull/520)

### Fixed

- When specifying `-skip-errors`, the batch spec would always be empty. This ix fixed and the intended behavior has been restored. [#539](https://github.com/sourcegraph/src-cli/pull/539)

## 3.27.0

### Added

- Extension publishing will now add a `gitHead` property to the extension's manifest. [#500](https://github.com/sourcegraph/src-cli/pull/500)
- `src batch [apply|preview]` now ignore repositories in which a `.batchignore` file exists. The `-force-override-ignore` flag can be used to turn that behaviour off. [#509](https://github.com/sourcegraph/src-cli/pull/509)
- `src search` now supports streaming search. If `src search` is called with the flag `-stream`, `src-cli` will stream back search results as they are found. In conjunction with `-stream` you can also specify `-display <integer>` to limit the number of results that will be displayed. For example, use `-display 0` if you are interested in the search statistics, such as number of results, but don't need to see the actual matches.

## 3.26.3

### Added

- Extension publishing will now add a `gitHead` property to the extension's manifest. [#500](https://github.com/sourcegraph/src-cli/pull/500)
- `src batch [apply|preview]` now ignore repositories in which a `.batchignore` file exists. The `-force-override-ignore` flag can be used to turn that behaviour off. [#509](https://github.com/sourcegraph/src-cli/pull/509)

## 3.26.2

### Fixed

- Publishing of Docker images for `src` was broken after release 3.24.3. This has been fixed, and [`sourcegraph/src-cli` images](https://hub.docker.com/r/sourcegraph/src-cli) are available once again. [#501](https://github.com/sourcegraph/src-cli/issues/501)

## 3.26.1

### Fixed

- Directly applying a batch change with `src batch apply` against Sourcegraph 3.25 or older would fail in 3.26.0. This has been fixed. [#495](https://github.com/sourcegraph/src-cli/issues/495)

## 3.26.0

### Added

- Two new [templating](https://docs.sourcegraph.com/campaigns/references/batch_spec_templating) variables have been added: `batch_change.name` and `batch_change.description`. [#491](https://github.com/sourcegraph/src-cli/pull/491)

### Changed

- Campaigns are now known as Batch Changes! The `src campaign` set of commands have been renamed to `src batch`; however, `src campaign` and `src campaigns` will be retained as aliases for `src batch` until the next major version of Sourcegraph. There should be no breaking changes as a result of this change. [#489](https://github.com/sourcegraph/src-cli/pull/489)

## 3.25.3

### Fixed

- The src login command now also properly respects the `-insecure-skip-verify` flag.

## 3.25.2

### Changed

- The volume workspace Docker image is now only pulled if the volume workspace mode is in use. [#477](https://github.com/sourcegraph/src-cli/pull/477)

### Fixed

- Using volume workspace mode could result in Git errors when used with Docker containers that do not run as root. These have been fixed. [#478](https://github.com/sourcegraph/src-cli/issues/478)

## 3.25.1

### Added

- Added a flag `-insecure-skip-verify` to disable TLS certificate validation.

### Changed

- Deprecated cache file formats are not read by `src campaign [apply|preview]` anymore.

## 3.25.0

### Added

- `src users delete` now asks for confirmation to delete all users when no user ID is provided. [#470](https://github.com/sourcegraph/src-cli/pull/470)

## 3.24.7

### Fixed

- The homebrew recipe for `src-cli` contained the wrong binary name.

## 3.24.6

### Fixed

- Workspaces could sometimes fail with docker bind mount errors, due to a race condition of multiple workspaces accessing the same auxilliary files. [#468](https://github.com/sourcegraph/src-cli/pull/468)

## 3.24.5

### Fixed

- Importing changesets was broken in the previous release and caused a SIGSEGV error.

## 3.24.4

### Added

- Experimental (requires Sourcegraph 3.25 or later): [`workspaces` in campaign specs](https://docs.sourcegraph.com/campaigns/references/campaign_spec_yaml_reference#workspaces) is now available to allow users to define multiple workspaces in a single repository. [#442](https://github.com/sourcegraph/src-cli/pull/442) and [#462](https://github.com/sourcegraph/src-cli/pull/462).
- The `changesetTemplate.published` field can now also be used to address a specific changeset in a repository by adding `@branch-of-changeset` at the end of the pattern. See [#461](https://github.com/sourcegraph/src-cli/pull/461) for an example and details.

### Fixed

- When `docker` becomes unresponsive `src campaign [apply|preview]` would get stuck and ignore Ctrl-C signals. That is now fixed.
- The `steps.files` attributes in campaign specs have been broken since 3.23.2 and now work again.

## 3.24.3

### Fixed

- 3.24.2 disabled the faster volume workspace on macOS when one or more non-root Docker images was used by a campaign to work around [the bug in issue #432](https://github.com/sourcegraph/src-cli/issues/432). This functionality has now been restored. [#434](https://github.com/sourcegraph/src-cli/pull/434)

## 3.24.2

### Fixed

- Executing campaigns on macOS 11 with Docker 3.1 could fail when using a volume workspace. This has been fixed. [#436](https://github.com/sourcegraph/src-cli/pull/436)

## 3.24.1

### Added

- Internal changes to the `src lsif upload` command. [#430](https://github.com/sourcegraph/src-cli/pull/430)

## 3.24.0

### Added

- `steps` in campaign specs can now have [`outputs`](https://docs.sourcegraph.com/campaigns/references/campaign_spec_yaml_reference#steps-outputs) that support [templating](https://docs.sourcegraph.com/campaigns/references/campaign_spec_templating). [#424](https://github.com/sourcegraph/src-cli/pull/424)
- `changesetTemplate` fields in campaign specs now also support [templating](https://docs.sourcegraph.com/campaigns/references/campaign_spec_templating). [#424](https://github.com/sourcegraph/src-cli/pull/424)

## 3.23.3

### Added

- Add verbosity flag to `lsif upload` action. Supply `-trace=1`, `-trace=2`, or `-trace=3` to the action to specify verbosity.

## 3.23.2

### Added

- `src campaign [apply|preview]` can now make use of Docker volumes, rather than bind-mounting the host filesystem. This is now the default on Intel macOS so long as the Docker images used in the campaign steps run as the same user, as volume mounts have generally better performance there. The optional `-workspace` flag can be used to override the default. [#412](https://github.com/sourcegraph/src-cli/pull/412)

### Changed

- `src login` now defaults to validating against `SRC_ENDPOINT` if configured.

### Fixed

- `src config` now works correctly when provided a subject.

## 3.23.1

### Fixed

- The src version command didn't send any authentication headers before, which could have failed for some instance configurations. The authentication header is now properly set for the request done in this command. [#411](https://github.com/sourcegraph/src-cli/pull/411)

## 3.23.0

### Added

- Experimental: [`transformChanges` in campaign specs](https://docs.sourcegraph.com/campaigns/references/campaign_spec_yaml_reference#transformchanges) is now available as a feature preview to allow users to create multiple changesets in a single repository. [#398](https://github.com/sourcegraph/src-cli/pull/398)

### Changed

- `src campaign [apply|preview]` now show the current execution progress in numbers next to the progress bar. [#396](https://github.com/sourcegraph/src-cli/pull/396)

### Fixed

- Two race conditions in the terminal UI of `src campaign [apply|preview]` have been fixed. [#399](https://github.com/sourcegraph/src-cli/pull/399)
- A regression caused repositories on unsupported code host to not be skipped by `src campaign [apply|preview]`, regardless of whether `-allow-unsupported` was set or not. [#403](https://github.com/sourcegraph/src-cli/pull/403)
- Previously `src campaign [apply|preview]` would crash when executing a campaign spec that contained `steps` but no `changesetTemplate`. [#406](https://github.com/sourcegraph/src-cli/issues/406)
- `src extensions copy` would copy an extension from the wrong Sourcegraph instance. [#409](https://github.com/sourcegraph/src-cli/pull/409)

### Removed

## 3.22.4

### Added

- Campaign steps may now include environment variables from outside of the campaign spec using [array syntax](http://docs.sourcegraph.com/campaigns/references/campaign_spec_yaml_reference#environment-array). [#392](https://github.com/sourcegraph/src-cli/pull/392)
- A new `-skip-errors` flag has been added to `src campaign [apply|preview]` to allow users to continue execution of and upload a campaign spec even if execution failed in some repositories. [#395](https://github.com/sourcegraph/src-cli/pull/395)

### Fixed

- The evaluation of the [`repository.branch` attribute](https://docs.sourcegraph.com/campaigns/references/campaign_spec_yaml_reference#on-repository) has been fixed to actually cause the correct version of the repository to be used. [#393](https://github.com/sourcegraph/src-cli/pull/393)
- Normally, when one or more repositories in a campaign generate an empty diff, a changeset spec isn't created. From src-cli 3.21.9 to 3.22.3, inclusive, re-running a campaign would result in an empty changeset spec being created by mistake if the empty changeset spec was in the execution cache, which would result in errors on Sourcegraph when applying the campaign. This has been fixed, and empty changeset specs in the cache are now treated the same way as uncached changeset specs that are empty: they are skipped, and a message is displayed in `-v` mode indicating the repo that was skipped. [#397](https://github.com/sourcegraph/src-cli/pull/397)

## 3.22.3

### Changed

- `src campaign [apply|preview]` now prints more detailed information about the diffs produced in each repository when run in verbose mode with `-v`. [#390](https://github.com/sourcegraph/src-cli/pull/390)
- The dependency `go-diff` has been upgraded to 0.6.1 to include https://github.com/sourcegraph/go-diff/pull/55.

## 3.22.2

### Fixed

- If `src campaign [validate|apply|preview]` was aborted while it was downloading repository archives it could leave behind partial ZIP files that would produce an error on the next run. This is now fixed by deleting partial files on abort. [#388](https://github.com/sourcegraph/src-cli/pull/388)
- A bug in `src campaign [apply|preview]` that would cause status bars in the TUI to not update has been fixed. [#389](https://github.com/sourcegraph/src-cli/pull/389)

## 3.22.1

### Fixed

- `src campaign [validate|apply|preview]` now print an error and usage information if a user accidentally provides an additional argument. [#384](https://github.com/sourcegraph/src-cli/pull/384)
- Fix a regression that was introduced by [#361](https://github.com/sourcegraph/src-cli/pull/361) and caused the "Resolving repositories" step of `src campaign [apply|preview]` to crash when the search query in the campaign spec yielded file matches and repository matches from the same repository.

## 3.22.0

### Fixed

- Fixed a bug that could cause `src campaign [apply|preview]` to crash in rare circumstances when executing a campaign spec due to a bug in the logic for the progress bar. [#378](https://github.com/sourcegraph/src-cli/pull/378)

## 3.21.9

### Added

- Commands for campaigns no longer require the `-namespace` parameter. If omitted, campaigns will use the currently authenticated user as the namespace. [#372](https://github.com/sourcegraph/src-cli/pull/372)
- `src campaign [apply|preview]` now caches the result of running steps in a repository even if they didn't produce changes.

## 3.21.8

### Fixed

- `src campaign [apply|preview]` could fail to parse the produced diff in a repository when `git` was configured to use a custom `diff` program. The fix is to ignore any local `git` configuration when running `git` commands. [#373](https://github.com/sourcegraph/src-cli/pull/373)

## 3.21.7

### Fixed

- Restored backward compatibility when creating campaigns against Sourcegraph 3.19, provided author details are not provided in the campaign spec. [#370](https://github.com/sourcegraph/src-cli/pull/370)

## 3.21.6

### Fixed

- Campaign steps run in a container that does not run as root could fail on systems that do not map the running user ID to the container, most notably desktop Linux. This has been fixed: temporary files and workspaces mounted into the container now have sufficient permissions to allow the container user to execute the step. [#366](https://github.com/sourcegraph/src-cli/pull/366)
- Executing campaigns on Windows would fail due to obscure `--cidfile` errors: namely, the temporary cidfile would not be removed before `docker run` was invoked. This has been fixed. [#368](https://github.com/sourcegraph/src-cli/pull/368)
- Unresponsive containers started by `src campaign [apply|preview]` can now be killed by hitting Ctrl-C. Previously the signal wasn't properly forwarded to the process in the container. [#369](https://github.com/sourcegraph/src-cli/pull/369)

## 3.21.5

### Added

- EXPERIMENTAL: Templated campaign specs and file mounting. The campaign specs evaluated by `src campaign [preview|apply]` can now include template variables in `steps.run`, `steps.env`, and the new `steps.files` property, which allows users to create files inside the container in which the step is executed. The feature is marked as EXPERIMENTAL because it might change in the near future until we deem it non-experimental. See [#361](https://github.com/sourcegraph/src-cli/pull/361) for details.

## 3.21.4

### Fixed

- The `src lsif upload` command now respects `SRC_HEADER_` environment variables for multipart uploads. These environment variables are described [here](AUTH_PROXY.md). [#360](https://github.com/sourcegraph/src-cli/pull/360)

## 3.21.3

### Changed

- The progress bar in `src campaign [preview|apply]` now shows when executing a step failed in a repository by styling the line red and displaying standard error output. [#355](https://github.com/sourcegraph/src-cli/pull/355)
- The `src lsif upload` command will give more informative output when an unexpected payload (non-JSON or non-unmarshallable) is received from the target endpoint. [#359](https://github.com/sourcegraph/src-cli/pull/359)

## 3.21.2

### Fixed

- The cache dir used by `src campaign [preview|apply]` is now created before trying to create files in it, fixing a bug where the first run of the command could fail with a "file doesn't exist" error message. [#352](https://github.com/sourcegraph/src-cli/pull/352)

## 3.21.1

### Added

- The `published` flag in campaign specs may now be an array, which allows only specific changesets within a campaign to be published based on the repository name. [#294](https://github.com/sourcegraph/src-cli/pull/294)
- A new `src campaign new` command creates a campaign spec YAML file with common values prefilled to make it easier to create a new campaign. [#339](https://github.com/sourcegraph/src-cli/pull/339)
- New experimental command [`src validate`](https://docs.sourcegraph.com/admin/validation) validates a Sourcegraph installation. [#200](https://github.com/sourcegraph/src-cli/pull/200)

### Changed

- Error reporting by `src campaign [preview|apply]` has been improved and now includes more information about which step failed in which repository. [#325](https://github.com/sourcegraph/src-cli/pull/325)
- The default behaviour of `src campaigns [preview|apply]` has been changed to retain downloaded archives of repositories for better performance across re-runs of the command. To use the old behaviour and delete the archives use the `-clean-archives` flag. Repository archives are also not stored in the directory for temp data (see `-tmp` flag) anymore but in the cache directory, which can be configured with the `-cache` flag. To manually delete archives between runs, delete the `*.zip` files in the `-cache` directory (see `src campaigns -help` for its default location).
- `src campaign [preview|apply]` now check whether `git` and `docker` are available before trying to execute a campaign spec's steps. [#326](https://github.com/sourcegraph/src-cli/pull/326)
- The progress bar displayed by `src campaign [preview|apply]` has been extended by status bars that show which steps are currently being executed for each repository. [#338](https://github.com/sourcegraph/src-cli/pull/338)
- `src campaign [preview|apply]` now shows a warning when no changeset specs have been created.
- Requests sent to Sourcegraph by the `src campaign` commands now use gzip compression for the body when talking to Sourcegraph 3.21.0 and later. [#336](https://github.com/sourcegraph/src-cli/pull/336) and [#343](https://github.com/sourcegraph/src-cli/pull/343)

### Fixed

- Log files created by `src campaigns [preview|apply]` are deleted again after successful execution. This was a regression and is not new behaviour. If steps failed to execute or the `-keep-logs` flag is set the log files are not cleaned up.
- `src campaign [preview|apply]` now correctly handle the interrupt signal (emitted in a terminal with Ctrl-C) and abort execution of campaign steps, cleaning up running Docker containers.

## 3.21.0

### Added

- The new `src login` subcommand helps you authenticate `src` to access your Sourcegraph instance with your user credentials. [#317](https://github.com/sourcegraph/src-cli/pull/312)

## 3.20.0

### Added

- Campaigns specs now include an optional `author` property. (If not included, `src campaigns` generates default values for author name and email.) `src campaigns` now includes the name and email in all changeset specs that it generates.
- The campaigns temp directory can now be overwritten by using the `-tmp` flag with `src campaigns [apply|preview]` or by setting `SRC_CAMPAIGNS_TMP_DIR`. The directory is used to, for example, store log files and unzipped repository archives when executing campaign specs.

### Changed

- Repositories without a default branch are skipped when applying/previewing a campaign spec. [#312](https://github.com/sourcegraph/src-cli/pull/312)
- Log files produced when applying/previewing a campaign spec now have the `.log` file extension for easier opening. [#315](https://github.com/sourcegraph/src-cli/pull/315)
- Campaign specs that apply to unsupported repositories will no longer generate an error. Instead, those repositories will be skipped by default and the campaign will be applied to the supported repositories only. [#314](https://github.com/sourcegraph/src-cli/pull/314)

### Fixed

- Empty changeset specs without a diff are no longer uploaded as part of a campaign spec. [#313](https://github.com/sourcegraph/src-cli/pull/313)

## 3.19.0

### Changed

- The default branch for the `src-cli` project has been changed to `main`. [#262](https://github.com/sourcegraph/src-cli/pull/262)

### Fixed

- `src campaigns` output has been improved in the Windows console. [#274](https://github.com/sourcegraph/src-cli/pull/274)
- `src campaigns` will no longer generate warnings if `user.name` or `user.email` have not been set in the global Git configuration. [#277](https://github.com/sourcegraph/src-cli/pull/277)

## 3.18.0

### Added

- Add `-dump-requests` as an option to all commands that interact with the Sourcegraph API. [#266](https://github.com/sourcegraph/src-cli/pull/266)

### Changed

- Reworked the `src campaigns` family of commands to [align with the new spec-based workflow](https://docs.sourcegraph.com/user/campaigns). Most notably, campaigns are now created and applied using the new `src campaigns apply` command, and use [the new YAML spec format](https://docs.sourcegraph.com/user/campaigns#creating-a-campaign). [#260](https://github.com/sourcegraph/src-cli/pull/260)

## 3.17.1

### Added

- Add -upload-route to the lsif upload command.

## 3.17.0

### Added

- New command `src serve-git` which can serve local repositories for Sourcegraph to clone. This was previously in a command called `src-expose`. See [serving local repositories](https://docs.sourcegraph.com/admin/external_service/src_serve_git) in our documentation to find out more. [#12363](https://github.com/sourcegraph/sourcegraph/issues/12363)
- When used with Sourcegraph 3.18 or later, campaigns can now be created on GitLab. [#231](https://github.com/sourcegraph/src-cli/pull/231)

## 3.16.1

### Fixed

- Fix inferred root for lsif upload command. [#248](https://github.com/sourcegraph/src-cli/pull/248)

### Removed

- Removed `clone-in-progress` flag. [#246](https://github.com/sourcegraph/src-cli/pull/246)

## 3.16

### Added

- Add `--no-progress` flag to the `lsif upload` command to disable verbose output in non-TTY environments.
- `SRC_HEADER_AUTHORIZATION="Bearer $(...)"` is now supported for authenticating `src` with custom auth proxies. See [auth proxy configuration docs](AUTH_PROXY.md) for more information. [#239](https://github.com/sourcegraph/src-cli/pull/239)
- Pull missing docker images automatically. [#191](https://github.com/sourcegraph/src-cli/pull/191)
- Searches that result in errors will now display any alerts returned by Sourcegraph, including suggestions for how the search could be corrected. [#221](https://github.com/sourcegraph/src-cli/pull/221)

### Changed

- The terminal UI has been replaced by the logger-based UI that was previously only visible in verbose-mode (`-v`). [#228](https://github.com/sourcegraph/src-cli/pull/228)
- Deprecated the `-endpoint` flag. Instead, use the `SRC_ENDPOINT` environment variable. [#235](https://github.com/sourcegraph/src-cli/pull/235)
