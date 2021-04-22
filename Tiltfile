print("Wallet KYC")

load("ext://restart_process", "docker_build_with_restart")

cfg = read_yaml(
    "tilt.yaml",
    default = read_yaml("tilt.yaml.sample"),
)

local_resource(
    "kyc-build-binary",
    "make fast_build",
    deps = ["./cmd", "./internal", "./service"],
)

docker_build(
    "velmie/wallet-kyc-db-migration",
    ".",
    dockerfile = "Dockerfile.migrations",
    only = "migrations",
)
k8s_resource(
    "wallet-kyc-db-migration",
    trigger_mode = TRIGGER_MODE_MANUAL,
    resource_deps = ["wallet-kyc-db-init"],
)

wallet_kyc_options = dict(
    entrypoint = "/app/service_kyc",
    dockerfile = "Dockerfile.prebuild",
    port_forwards = [],
    helm_set = [],
)

if cfg["debug"]:
    wallet_kyc_options["entrypoint"] = "$GOPATH/bin/dlv --continue --listen :%s --accept-multiclient --api-version=2 --headless=true exec /app/service_kyc" % cfg["debug_port"]
    wallet_kyc_options["dockerfile"] = "Dockerfile.debug"
    wallet_kyc_options["port_forwards"] = cfg["debug_port"]
    wallet_kyc_options["helm_set"] = ["containerLivenessProbe.enabled=false", "containerPorts[0].containerPort=%s" % cfg["debug_port"]]

docker_build_with_restart(
    "velmie/wallet-kyc",
    ".",
    dockerfile = wallet_kyc_options["dockerfile"],
    entrypoint = wallet_kyc_options["entrypoint"],
    only = [
        "./build",
        "zoneinfo.zip",
    ],
    live_update = [
        sync("./build", "/app/"),
    ],
)
k8s_resource(
    "wallet-kyc",
    resource_deps = ["wallet-kyc-db-migration"],
    port_forwards = wallet_kyc_options["port_forwards"],
)

yaml = helm(
    "./helm/wallet-kyc",
    # The release name, equivalent to helm --name
    name = "wallet-kyc",
    # The values file to substitute into the chart.
    values = ["./helm/values-dev.yaml"],
    set = wallet_kyc_options["helm_set"],
)

k8s_yaml(yaml)
