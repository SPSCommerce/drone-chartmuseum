# Drone plugin for ChartMuseum

Drone plugin to package and upload Helm charts to [ChartMuseum](https://github.com/kubernetes-helm/chartmuseum)

The SPS Chartmuseum url is (https://chartmuseum.spsapps.net)[https://chartmuseum.spsapps.net] and implemented in git repo (sps-chartmuseum)[https://github.com/SPSCommerce/sps-chartmuseum]

When managing Charts for your organization, you may either choose to put Chart definitions within each project or centralised in a `helm-charts` repository. The official public-charts repo is an example of the latter.

This plugin supports both approaches as well as the ability to detect and process only changes as part of a git repository.

## Usage Examples

- Process all charts from root of repository

  Package all charts under `chart_dir` and upload to Repository server.

  ```YAML
  pipeline:
    chartmuseum-all:
      image: spscommerce/plugin-helmchart-publish
      when:
        branch: [master]
  ```

- Process only changed charts

  Detect changed files between `previous_commit` and `current_commit`, only package and upload modified helm charts. Ignores modifications if they match `.helmignore` rules.

  ```YAML
  pipeline:
    chartmuseum-diff:
      image: spscommerce/plugin-helmchart-publish
      previous_commit: ${DRONE_PREV_COMMIT_SHA}
      current_commit: ${DRONE_COMMIT_SHA}
      when:
        branch: [master]
  ```

- Process only a specific chart. Can be combined with commit SHA to only process if chart is modified. (also uses `.helmignore`)

  ```YAML
  pipeline:
    chartmuseum-single:
      image: spscommerce/plugin-helmchart-publish
      chart_path: nginx
      when:
        branch: [master]
  ```


## Full utilisation details

```bash
NAME:
   drone-chartmuseum-plugin - drone plugin to upload charts to chartmuseum server

USAGE:
   drone-chartmuseum [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --repo-url value, -u value                   ChartMuseum API base URL [$PLUGIN_REPO_URL, $REPO_URL]
   --username value, -n value                   Chartmuseum HTTP Basic auth username [$PLUGIN_REPO_USERNAME]
   --password value, -x value                   Chartmuseum HTTP Basic auth password [$PLUGIN_REPO_PASSWORD]
   --chart-path value, -i value                 Path to chart, relative to charts-dir [$PLUGIN_CHART_PATH, $CHART_PATH]
   --charts-dir value, -d value                 chart directory (default: "./") [$PLUGIN_CHARTS_DIR, $CHARTS_DIR]
   --save-dir value, -o value                   Directory to save chart packages (default: "uploads/") [$PLUGIN_SAVE_DIR, $SAVE_DIR]
   --previous-commit COMMIT_SHA, -p COMMIT_SHA  Previous commit id (COMMIT_SHA) [$PLUGIN_PREVIOUS_COMMIT, $PREVIOUS_COMMIT]
   --current-commit COMMIT_SHA, -c COMMIT_SHA   Current commit id (COMMIT_SHA) [$PLUGIN_CURRENT_COMMIT, $CURRENT_COMMIT]
   --log-level value                            Log level (panic, fatal, error, warn, info, or debug) (default: "error") [$PLUGIN_LOG_LEVEL, $LOG_LEVEL]
   --help, -h                                   show help
   --version, -v                                print the version
```

```bash
docker run --rm \
  -e PLUGIN_REPO_URL="http://helm-charts.example.com" \
  -e PLUGIN_PREVIOUS_COMMIT="<commit-sha>" \
  -e PLUGIN_CURRENT_COMMIT="<commit-sha>" \
  quay.io/honestbee/drone-chartmuseum
```

## Unit Tests

Unit tests support log level also, though you may need to clean cache when changing log level.

```bash
go clean -cache
LOG_LEVEL=debug go test -v ./...
```

## To Do

- Support chart dependencies
