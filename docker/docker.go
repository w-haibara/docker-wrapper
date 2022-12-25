package docker

import (
	"fmt"
	"os/exec"
)

type DockerOption struct {
	/*
		Location of client config files
	*/
	Config *string

	/*
		Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with "docker context use")
	*/
	Context *string

	/*
		Enable debug mode
	*/
	Debug *bool

	/*
		Daemon socket(s) to connect to
	*/
	Host []string

	/*
		Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")
	*/
	LogLevel *string

	/*
		Use TLS; implied by --tlsverify
	*/
	Tls *bool

	/*
		Trust certs signed only by this CA
	*/
	Tlscacert *string

	/*
		Path to TLS certificate file
	*/
	Tlscert *string

	/*
		Path to TLS key file
	*/
	Tlskey *string

	/*
		Use TLS and verify the remote
	*/
	Tlsverify *bool

	/*
		Print version information and quit
	*/
	Version *bool
}

/*
DockerCmd is wrapper of 'docker'
------------------------------
docker [OPTIONS] COMMAND [ARG...]
A self-sufficient runtime for containers
------------------------------
*/
func DockerCmd(opt DockerOption, args []string) *exec.Cmd {
	cargs := []string{}
	if opt.Config != nil {
		cargs = append(cargs, "--config="+fmt.Sprint(*opt.Config))
	}
	if opt.Context != nil {
		cargs = append(cargs, "--context="+fmt.Sprint(*opt.Context))
	}
	if opt.Debug != nil {
		cargs = append(cargs, "--debug="+fmt.Sprint(*opt.Debug))
	}
	if opt.Host != nil {
		for _, str := range opt.Host {
			cargs = append(cargs, "--host")
			cargs = append(cargs, str)
		}
	}
	if opt.LogLevel != nil {
		cargs = append(cargs, "--log-level="+fmt.Sprint(*opt.LogLevel))
	}
	if opt.Tls != nil {
		cargs = append(cargs, "--tls="+fmt.Sprint(*opt.Tls))
	}
	if opt.Tlscacert != nil {
		cargs = append(cargs, "--tlscacert="+fmt.Sprint(*opt.Tlscacert))
	}
	if opt.Tlscert != nil {
		cargs = append(cargs, "--tlscert="+fmt.Sprint(*opt.Tlscert))
	}
	if opt.Tlskey != nil {
		cargs = append(cargs, "--tlskey="+fmt.Sprint(*opt.Tlskey))
	}
	if opt.Tlsverify != nil {
		cargs = append(cargs, "--tlsverify="+fmt.Sprint(*opt.Tlsverify))
	}
	if opt.Version != nil {
		cargs = append(cargs, "--version="+fmt.Sprint(*opt.Version))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerAttachOption struct {
	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Do not attach STDIN
	*/
	NoStdin *bool

	/*
		Proxy all received signals to the process
	*/
	SigProxy *bool
}

/*
DockerAttachCmd is wrapper of 'docker attach'
------------------------------
attach [OPTIONS] CONTAINER
Attach local standard input, output, and error streams to a running container
------------------------------
*/
func DockerAttachCmd(opt DockerAttachOption, args []string) *exec.Cmd {
	cargs := []string{"attach"}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.NoStdin != nil {
		cargs = append(cargs, "--no-stdin="+fmt.Sprint(*opt.NoStdin))
	}
	if opt.SigProxy != nil {
		cargs = append(cargs, "--sig-proxy="+fmt.Sprint(*opt.SigProxy))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerBuildOption struct {
	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	AddHost []string

	/*
		Set build-time variables
	*/
	BuildArg []string

	/*
		Images to consider as cache sources
	*/
	CacheFrom *string

	/*
		Optional parent cgroup for the container
	*/
	CgroupParent *string

	/*
		Compress the build context using gzip
	*/
	Compress *bool

	/*
		Limit the CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit the CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Name of the Dockerfile (Default is 'PATH/Dockerfile')
	*/
	File *string

	/*
		Always remove intermediate containers
	*/
	ForceRm *bool

	/*
		Write the image ID to the file
	*/
	Iidfile *string

	/*
		Container isolation technology
	*/
	Isolation *string

	/*
		Set metadata for an image
	*/
	Label []string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Set the networking mode for the RUN instructions during build
	*/
	Network *string

	/*
		Do not use cache when building the image
	*/
	NoCache *bool

	/*
		Output destination (format: type=local,dest=path)
	*/
	Output *string

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Set type of progress output (auto, plain, tty). Use plain to show container output
	*/
	Progress *string

	/*
		Always attempt to pull a newer version of the image
	*/
	Pull *bool

	/*
		Suppress the build output and print image ID on success
	*/
	Quiet *bool

	/*
		Remove intermediate containers after a successful build
	*/
	Rm *bool

	/*
		Secret file to expose to the build (only if BuildKit enabled): id=mysecret,src=/local/secret
	*/
	Secret *string

	/*
		Security options
	*/
	SecurityOpt *string

	/*
		Size of /dev/shm
	*/
	ShmSize *string

	/*
		Squash newly built layers into a single new layer
	*/
	Squash *bool

	/*
		SSH agent socket or keys to expose to the build (only if BuildKit enabled) (format: default|<id>[=<socket>|<key>[,<key>]])
	*/
	Ssh *string

	/*
		Stream attaches to server to negotiate build context
	*/
	Stream *bool

	/*
		Name and optionally a tag in the 'name:tag' format
	*/
	Tag []string

	/*
		Set the target build stage to build.
	*/
	Target *string

	/*
		Ulimit options
	*/
	Ulimit *string
}

/*
DockerBuildCmd is wrapper of 'docker build'
------------------------------
build [OPTIONS] PATH | URL | -
Build an image from a Dockerfile
------------------------------
*/
func DockerBuildCmd(opt DockerBuildOption, args []string) *exec.Cmd {
	cargs := []string{"build"}
	if opt.AddHost != nil {
		for _, str := range opt.AddHost {
			cargs = append(cargs, "--add-host")
			cargs = append(cargs, str)
		}
	}
	if opt.BuildArg != nil {
		for _, str := range opt.BuildArg {
			cargs = append(cargs, "--build-arg")
			cargs = append(cargs, str)
		}
	}
	if opt.CacheFrom != nil {
		cargs = append(cargs, "--cache-from="+fmt.Sprint(*opt.CacheFrom))
	}
	if opt.CgroupParent != nil {
		cargs = append(cargs, "--cgroup-parent="+fmt.Sprint(*opt.CgroupParent))
	}
	if opt.Compress != nil {
		cargs = append(cargs, "--compress="+fmt.Sprint(*opt.Compress))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.File != nil {
		cargs = append(cargs, "--file="+fmt.Sprint(*opt.File))
	}
	if opt.ForceRm != nil {
		cargs = append(cargs, "--force-rm="+fmt.Sprint(*opt.ForceRm))
	}
	if opt.Iidfile != nil {
		cargs = append(cargs, "--iidfile="+fmt.Sprint(*opt.Iidfile))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NoCache != nil {
		cargs = append(cargs, "--no-cache="+fmt.Sprint(*opt.NoCache))
	}
	if opt.Output != nil {
		cargs = append(cargs, "--output="+fmt.Sprint(*opt.Output))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Progress != nil {
		cargs = append(cargs, "--progress="+fmt.Sprint(*opt.Progress))
	}
	if opt.Pull != nil {
		cargs = append(cargs, "--pull="+fmt.Sprint(*opt.Pull))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Rm != nil {
		cargs = append(cargs, "--rm="+fmt.Sprint(*opt.Rm))
	}
	if opt.Secret != nil {
		cargs = append(cargs, "--secret="+fmt.Sprint(*opt.Secret))
	}
	if opt.SecurityOpt != nil {
		cargs = append(cargs, "--security-opt="+fmt.Sprint(*opt.SecurityOpt))
	}
	if opt.ShmSize != nil {
		cargs = append(cargs, "--shm-size="+fmt.Sprint(*opt.ShmSize))
	}
	if opt.Squash != nil {
		cargs = append(cargs, "--squash="+fmt.Sprint(*opt.Squash))
	}
	if opt.Ssh != nil {
		cargs = append(cargs, "--ssh="+fmt.Sprint(*opt.Ssh))
	}
	if opt.Stream != nil {
		cargs = append(cargs, "--stream="+fmt.Sprint(*opt.Stream))
	}
	if opt.Tag != nil {
		for _, str := range opt.Tag {
			cargs = append(cargs, "--tag")
			cargs = append(cargs, str)
		}
	}
	if opt.Target != nil {
		cargs = append(cargs, "--target="+fmt.Sprint(*opt.Target))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerBuilderCmd is wrapper of 'docker builder'
------------------------------
builder
Manage builds
------------------------------
*/
func DockerBuilderCmd(args []string) *exec.Cmd {
	cargs := []string{"builder"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerBuilderBuildOption struct {
	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	AddHost []string

	/*
		Set build-time variables
	*/
	BuildArg []string

	/*
		Images to consider as cache sources
	*/
	CacheFrom *string

	/*
		Optional parent cgroup for the container
	*/
	CgroupParent *string

	/*
		Compress the build context using gzip
	*/
	Compress *bool

	/*
		Limit the CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit the CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Name of the Dockerfile (Default is 'PATH/Dockerfile')
	*/
	File *string

	/*
		Always remove intermediate containers
	*/
	ForceRm *bool

	/*
		Write the image ID to the file
	*/
	Iidfile *string

	/*
		Container isolation technology
	*/
	Isolation *string

	/*
		Set metadata for an image
	*/
	Label []string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Set the networking mode for the RUN instructions during build
	*/
	Network *string

	/*
		Do not use cache when building the image
	*/
	NoCache *bool

	/*
		Output destination (format: type=local,dest=path)
	*/
	Output *string

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Set type of progress output (auto, plain, tty). Use plain to show container output
	*/
	Progress *string

	/*
		Always attempt to pull a newer version of the image
	*/
	Pull *bool

	/*
		Suppress the build output and print image ID on success
	*/
	Quiet *bool

	/*
		Remove intermediate containers after a successful build
	*/
	Rm *bool

	/*
		Secret file to expose to the build (only if BuildKit enabled): id=mysecret,src=/local/secret
	*/
	Secret *string

	/*
		Security options
	*/
	SecurityOpt *string

	/*
		Size of /dev/shm
	*/
	ShmSize *string

	/*
		Squash newly built layers into a single new layer
	*/
	Squash *bool

	/*
		SSH agent socket or keys to expose to the build (only if BuildKit enabled) (format: default|<id>[=<socket>|<key>[,<key>]])
	*/
	Ssh *string

	/*
		Stream attaches to server to negotiate build context
	*/
	Stream *bool

	/*
		Name and optionally a tag in the 'name:tag' format
	*/
	Tag []string

	/*
		Set the target build stage to build.
	*/
	Target *string

	/*
		Ulimit options
	*/
	Ulimit *string
}

/*
DockerBuilderBuildCmd is wrapper of 'docker builder build'
------------------------------
build [OPTIONS] PATH | URL | -
Build an image from a Dockerfile
------------------------------
*/
func DockerBuilderBuildCmd(opt DockerBuilderBuildOption, args []string) *exec.Cmd {
	cargs := []string{"builder", "build"}
	if opt.AddHost != nil {
		for _, str := range opt.AddHost {
			cargs = append(cargs, "--add-host")
			cargs = append(cargs, str)
		}
	}
	if opt.BuildArg != nil {
		for _, str := range opt.BuildArg {
			cargs = append(cargs, "--build-arg")
			cargs = append(cargs, str)
		}
	}
	if opt.CacheFrom != nil {
		cargs = append(cargs, "--cache-from="+fmt.Sprint(*opt.CacheFrom))
	}
	if opt.CgroupParent != nil {
		cargs = append(cargs, "--cgroup-parent="+fmt.Sprint(*opt.CgroupParent))
	}
	if opt.Compress != nil {
		cargs = append(cargs, "--compress="+fmt.Sprint(*opt.Compress))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.File != nil {
		cargs = append(cargs, "--file="+fmt.Sprint(*opt.File))
	}
	if opt.ForceRm != nil {
		cargs = append(cargs, "--force-rm="+fmt.Sprint(*opt.ForceRm))
	}
	if opt.Iidfile != nil {
		cargs = append(cargs, "--iidfile="+fmt.Sprint(*opt.Iidfile))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NoCache != nil {
		cargs = append(cargs, "--no-cache="+fmt.Sprint(*opt.NoCache))
	}
	if opt.Output != nil {
		cargs = append(cargs, "--output="+fmt.Sprint(*opt.Output))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Progress != nil {
		cargs = append(cargs, "--progress="+fmt.Sprint(*opt.Progress))
	}
	if opt.Pull != nil {
		cargs = append(cargs, "--pull="+fmt.Sprint(*opt.Pull))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Rm != nil {
		cargs = append(cargs, "--rm="+fmt.Sprint(*opt.Rm))
	}
	if opt.Secret != nil {
		cargs = append(cargs, "--secret="+fmt.Sprint(*opt.Secret))
	}
	if opt.SecurityOpt != nil {
		cargs = append(cargs, "--security-opt="+fmt.Sprint(*opt.SecurityOpt))
	}
	if opt.ShmSize != nil {
		cargs = append(cargs, "--shm-size="+fmt.Sprint(*opt.ShmSize))
	}
	if opt.Squash != nil {
		cargs = append(cargs, "--squash="+fmt.Sprint(*opt.Squash))
	}
	if opt.Ssh != nil {
		cargs = append(cargs, "--ssh="+fmt.Sprint(*opt.Ssh))
	}
	if opt.Stream != nil {
		cargs = append(cargs, "--stream="+fmt.Sprint(*opt.Stream))
	}
	if opt.Tag != nil {
		for _, str := range opt.Tag {
			cargs = append(cargs, "--tag")
			cargs = append(cargs, str)
		}
	}
	if opt.Target != nil {
		cargs = append(cargs, "--target="+fmt.Sprint(*opt.Target))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerBuilderPruneOption struct {
	/*
		Remove all unused build cache, not just dangling ones
	*/
	All *bool

	/*
		Provide filter values (e.g. 'until=24h')
	*/
	Filter *string

	/*
		Do not prompt for confirmation
	*/
	Force *bool

	/*
		Amount of disk space to keep for cache
	*/
	KeepStorage *string
}

/*
DockerBuilderPruneCmd is wrapper of 'docker builder prune'
------------------------------
prune
Remove build cache
------------------------------
*/
func DockerBuilderPruneCmd(opt DockerBuilderPruneOption, args []string) *exec.Cmd {
	cargs := []string{"builder", "prune"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	if opt.KeepStorage != nil {
		cargs = append(cargs, "--keep-storage="+fmt.Sprint(*opt.KeepStorage))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerCheckpointCmd is wrapper of 'docker checkpoint'
------------------------------
checkpoint
Manage checkpoints
------------------------------
*/
func DockerCheckpointCmd(args []string) *exec.Cmd {
	cargs := []string{"checkpoint"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerCheckpointCreateOption struct {
	/*
		Use a custom checkpoint storage directory
	*/
	CheckpointDir *string

	/*
		Leave the container running after checkpoint
	*/
	LeaveRunning *bool
}

/*
DockerCheckpointCreateCmd is wrapper of 'docker checkpoint create'
------------------------------
create [OPTIONS] CONTAINER CHECKPOINT
Create a checkpoint from a running container
------------------------------
*/
func DockerCheckpointCreateCmd(opt DockerCheckpointCreateOption, args []string) *exec.Cmd {
	cargs := []string{"checkpoint", "create"}
	if opt.CheckpointDir != nil {
		cargs = append(cargs, "--checkpoint-dir="+fmt.Sprint(*opt.CheckpointDir))
	}
	if opt.LeaveRunning != nil {
		cargs = append(cargs, "--leave-running="+fmt.Sprint(*opt.LeaveRunning))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerCheckpointLsOption struct {
	/*
		Use a custom checkpoint storage directory
	*/
	CheckpointDir *string
}

/*
DockerCheckpointLsCmd is wrapper of 'docker checkpoint ls'
------------------------------
ls [OPTIONS] CONTAINER
List checkpoints for a container
------------------------------
*/
func DockerCheckpointLsCmd(opt DockerCheckpointLsOption, args []string) *exec.Cmd {
	cargs := []string{"checkpoint", "ls"}
	if opt.CheckpointDir != nil {
		cargs = append(cargs, "--checkpoint-dir="+fmt.Sprint(*opt.CheckpointDir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerCheckpointRmOption struct {
	/*
		Use a custom checkpoint storage directory
	*/
	CheckpointDir *string
}

/*
DockerCheckpointRmCmd is wrapper of 'docker checkpoint rm'
------------------------------
rm [OPTIONS] CONTAINER CHECKPOINT
Remove a checkpoint
------------------------------
*/
func DockerCheckpointRmCmd(opt DockerCheckpointRmOption, args []string) *exec.Cmd {
	cargs := []string{"checkpoint", "rm"}
	if opt.CheckpointDir != nil {
		cargs = append(cargs, "--checkpoint-dir="+fmt.Sprint(*opt.CheckpointDir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerCommitOption struct {
	/*
		Author (e.g., "John Hannibal Smith <hannibal@a-team.com>")
	*/
	Author *string

	/*
		Apply Dockerfile instruction to the created image
	*/
	Change []string

	/*
		Commit message
	*/
	Message *string

	/*
		Pause container during commit
	*/
	Pause *bool
}

/*
DockerCommitCmd is wrapper of 'docker commit'
------------------------------
commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]
Create a new image from a container's changes
------------------------------
*/
func DockerCommitCmd(opt DockerCommitOption, args []string) *exec.Cmd {
	cargs := []string{"commit"}
	if opt.Author != nil {
		cargs = append(cargs, "--author="+fmt.Sprint(*opt.Author))
	}
	if opt.Change != nil {
		for _, str := range opt.Change {
			cargs = append(cargs, "--change")
			cargs = append(cargs, str)
		}
	}
	if opt.Message != nil {
		cargs = append(cargs, "--message="+fmt.Sprint(*opt.Message))
	}
	if opt.Pause != nil {
		cargs = append(cargs, "--pause="+fmt.Sprint(*opt.Pause))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerConfigCmd is wrapper of 'docker config'
------------------------------
config
Manage Docker configs
------------------------------
*/
func DockerConfigCmd(args []string) *exec.Cmd {
	cargs := []string{"config"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerConfigCreateOption struct {
	/*
		Config labels
	*/
	Label []string

	/*
		Template driver
	*/
	TemplateDriver *string
}

/*
DockerConfigCreateCmd is wrapper of 'docker config create'
------------------------------
create [OPTIONS] CONFIG file|-
Create a config from a file or STDIN
------------------------------
*/
func DockerConfigCreateCmd(opt DockerConfigCreateOption, args []string) *exec.Cmd {
	cargs := []string{"config", "create"}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.TemplateDriver != nil {
		cargs = append(cargs, "--template-driver="+fmt.Sprint(*opt.TemplateDriver))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerConfigInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Print the information in a human friendly format
	*/
	Pretty *bool
}

/*
DockerConfigInspectCmd is wrapper of 'docker config inspect'
------------------------------
inspect [OPTIONS] CONFIG [CONFIG...]
Display detailed information on one or more configs
------------------------------
*/
func DockerConfigInspectCmd(opt DockerConfigInspectOption, args []string) *exec.Cmd {
	cargs := []string{"config", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Pretty != nil {
		cargs = append(cargs, "--pretty="+fmt.Sprint(*opt.Pretty))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerConfigLsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print configs using a Go template
	*/
	Format *string

	/*
		Only display IDs
	*/
	Quiet *bool
}

/*
DockerConfigLsCmd is wrapper of 'docker config ls'
------------------------------
ls [OPTIONS]
List configs
------------------------------
*/
func DockerConfigLsCmd(opt DockerConfigLsOption, args []string) *exec.Cmd {
	cargs := []string{"config", "ls"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerConfigRmCmd is wrapper of 'docker config rm'
------------------------------
rm CONFIG [CONFIG...]
Remove one or more configs
------------------------------
*/
func DockerConfigRmCmd(args []string) *exec.Cmd {
	cargs := []string{"config", "rm"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerCmd is wrapper of 'docker container'
------------------------------
container
Manage containers
------------------------------
*/
func DockerContainerCmd(args []string) *exec.Cmd {
	cargs := []string{"container"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerAttachOption struct {
	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Do not attach STDIN
	*/
	NoStdin *bool

	/*
		Proxy all received signals to the process
	*/
	SigProxy *bool
}

/*
DockerContainerAttachCmd is wrapper of 'docker container attach'
------------------------------
attach [OPTIONS] CONTAINER
Attach local standard input, output, and error streams to a running container
------------------------------
*/
func DockerContainerAttachCmd(opt DockerContainerAttachOption, args []string) *exec.Cmd {
	cargs := []string{"container", "attach"}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.NoStdin != nil {
		cargs = append(cargs, "--no-stdin="+fmt.Sprint(*opt.NoStdin))
	}
	if opt.SigProxy != nil {
		cargs = append(cargs, "--sig-proxy="+fmt.Sprint(*opt.SigProxy))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerCommitOption struct {
	/*
		Author (e.g., "John Hannibal Smith <hannibal@a-team.com>")
	*/
	Author *string

	/*
		Apply Dockerfile instruction to the created image
	*/
	Change []string

	/*
		Commit message
	*/
	Message *string

	/*
		Pause container during commit
	*/
	Pause *bool
}

/*
DockerContainerCommitCmd is wrapper of 'docker container commit'
------------------------------
commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]
Create a new image from a container's changes
------------------------------
*/
func DockerContainerCommitCmd(opt DockerContainerCommitOption, args []string) *exec.Cmd {
	cargs := []string{"container", "commit"}
	if opt.Author != nil {
		cargs = append(cargs, "--author="+fmt.Sprint(*opt.Author))
	}
	if opt.Change != nil {
		for _, str := range opt.Change {
			cargs = append(cargs, "--change")
			cargs = append(cargs, str)
		}
	}
	if opt.Message != nil {
		cargs = append(cargs, "--message="+fmt.Sprint(*opt.Message))
	}
	if opt.Pause != nil {
		cargs = append(cargs, "--pause="+fmt.Sprint(*opt.Pause))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerCpOption struct {
	/*
		Archive mode (copy all uid/gid information)
	*/
	Archive *bool

	/*
		Always follow symbol link in SRC_PATH
	*/
	FollowLink *bool
}

/*
DockerContainerCpCmd is wrapper of 'docker container cp'
------------------------------
cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-

	docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH

Copy files/folders between a container and the local filesystem
------------------------------
*/
func DockerContainerCpCmd(opt DockerContainerCpOption, args []string) *exec.Cmd {
	cargs := []string{"container", "cp"}
	if opt.Archive != nil {
		cargs = append(cargs, "--archive="+fmt.Sprint(*opt.Archive))
	}
	if opt.FollowLink != nil {
		cargs = append(cargs, "--follow-link="+fmt.Sprint(*opt.FollowLink))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerCreateOption struct {
	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	AddHost []string

	/*
		Attach to STDIN, STDOUT or STDERR
	*/
	Attach []string

	/*
		Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
	*/
	BlkioWeight *uint16

	/*
		Block IO weight (relative device weight)
	*/
	BlkioWeightDevice []string

	/*
		Add Linux capabilities
	*/
	CapAdd []string

	/*
		Drop Linux capabilities
	*/
	CapDrop []string

	/*
		Optional parent cgroup for the container
	*/
	CgroupParent *string

	/*
			Cgroup namespace to use (host|private)
		'host':    Run the container in the Docker host's cgroup namespace
		'private': Run the container in its own private cgroup namespace
		'':        Use the cgroup namespace as configured by the
		           default-cgroupns-mode option on the daemon (default)
	*/
	Cgroupns *string

	/*
		Write the container ID to the file
	*/
	Cidfile *string

	/*
		CPU count (Windows only)
	*/
	CpuCount *int64

	/*
		CPU percent (Windows only)
	*/
	CpuPercent *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		Limit CPU real-time period in microseconds
	*/
	CpuRtPeriod *int64

	/*
		Limit CPU real-time runtime in microseconds
	*/
	CpuRtRuntime *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		Number of CPUs
	*/
	Cpus *string

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Add a host device to the container
	*/
	Device []string

	/*
		Add a rule to the cgroup allowed devices list
	*/
	DeviceCgroupRule []string

	/*
		Limit read rate (bytes per second) from a device
	*/
	DeviceReadBps []string

	/*
		Limit read rate (IO per second) from a device
	*/
	DeviceReadIops []string

	/*
		Limit write rate (bytes per second) to a device
	*/
	DeviceWriteBps []string

	/*
		Limit write rate (IO per second) to a device
	*/
	DeviceWriteIops []string

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Set custom DNS servers
	*/
	Dns []string

	/*
		Set DNS options
	*/
	DnsOpt []string

	/*
		Set DNS options
	*/
	DnsOption []string

	/*
		Set custom DNS search domains
	*/
	DnsSearch []string

	/*
		Container NIS domain name
	*/
	Domainname *string

	/*
		Overwrite the default ENTRYPOINT of the image
	*/
	Entrypoint *string

	/*
		Set environment variables
	*/
	Env []string

	/*
		Read in a file of environment variables
	*/
	EnvFile []string

	/*
		Expose a port or a range of ports
	*/
	Expose []string

	/*
		GPU devices to add to the container ('all' to pass all GPUs)
	*/
	Gpus *string

	/*
		Add additional groups to join
	*/
	GroupAdd []string

	/*
		Command to run to check health
	*/
	HealthCmd *string

	/*
		Time between running the check (ms|s|m|h) (default 0s)
	*/
	HealthInterval *string

	/*
		Consecutive failures needed to report unhealthy
	*/
	HealthRetries *int

	/*
		Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
	*/
	HealthStartPeriod *string

	/*
		Maximum time to allow one check to run (ms|s|m|h) (default 0s)
	*/
	HealthTimeout *string

	/*
		Print usage
	*/
	Help *bool

	/*
		Container host name
	*/
	Hostname *string

	/*
		Run an init inside the container that forwards signals and reaps processes
	*/
	Init *bool

	/*
		Keep STDIN open even if not attached
	*/
	Interactive *bool

	/*
		Maximum IO bandwidth limit for the system drive (Windows only)
	*/
	IoMaxbandwidth *string

	/*
		Maximum IOps limit for the system drive (Windows only)
	*/
	IoMaxiops *uint64

	/*
		IPv4 address (e.g., 172.30.100.104)
	*/
	Ip *string

	/*
		IPv6 address (e.g., 2001:db8::33)
	*/
	Ip6 *string

	/*
		IPC mode to use
	*/
	Ipc *string

	/*
		Container isolation technology
	*/
	Isolation *string

	/*
		Kernel memory limit
	*/
	KernelMemory *string

	/*
		Set meta data on a container
	*/
	Label []string

	/*
		Read in a line delimited file of labels
	*/
	LabelFile []string

	/*
		Add link to another container
	*/
	Link []string

	/*
		Container IPv4/IPv6 link-local addresses
	*/
	LinkLocalIp []string

	/*
		Logging driver for the container
	*/
	LogDriver *string

	/*
		Log driver options
	*/
	LogOpt []string

	/*
		Container MAC address (e.g., 92:d0:c6:0a:29:33)
	*/
	MacAddress *string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Memory soft limit
	*/
	MemoryReservation *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Tune container memory swappiness (0 to 100)
	*/
	MemorySwappiness *int64

	/*
		Attach a filesystem mount to the container
	*/
	Mount *string

	/*
		Assign a name to the container
	*/
	Name *string

	/*
		Connect a container to a network
	*/
	Net *string

	/*
		Add network-scoped alias for the container
	*/
	NetAlias []string

	/*
		Connect a container to a network
	*/
	Network *string

	/*
		Add network-scoped alias for the container
	*/
	NetworkAlias []string

	/*
		Disable any container-specified HEALTHCHECK
	*/
	NoHealthcheck *bool

	/*
		Disable OOM Killer
	*/
	OomKillDisable *bool

	/*
		Tune host's OOM preferences (-1000 to 1000)
	*/
	OomScoreAdj *int

	/*
		PID namespace to use
	*/
	Pid *string

	/*
		Tune container pids limit (set -1 for unlimited)
	*/
	PidsLimit *int64

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Give extended privileges to this container
	*/
	Privileged *bool

	/*
		Publish a container's port(s) to the host
	*/
	Publish []string

	/*
		Publish all exposed ports to random ports
	*/
	PublishAll *bool

	/*
		Pull image before creating ("always"|"missing"|"never")
	*/
	Pull *string

	/*
		Mount the container's root filesystem as read only
	*/
	ReadOnly *bool

	/*
		Restart policy to apply when a container exits
	*/
	Restart *string

	/*
		Automatically remove the container when it exits
	*/
	Rm *bool

	/*
		Runtime to use for this container
	*/
	Runtime *string

	/*
		Security Options
	*/
	SecurityOpt []string

	/*
		Size of /dev/shm
	*/
	ShmSize *string

	/*
		Signal to stop a container
	*/
	StopSignal *string

	/*
		Timeout (in seconds) to stop a container
	*/
	StopTimeout *int

	/*
		Storage driver options for the container
	*/
	StorageOpt []string

	/*
		Sysctl options
	*/
	Sysctl map[string]string

	/*
		Mount a tmpfs directory
	*/
	Tmpfs []string

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Ulimit options
	*/
	Ulimit *string

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		User namespace to use
	*/
	Userns *string

	/*
		UTS namespace to use
	*/
	Uts *string

	/*
		Bind mount a volume
	*/
	Volume []string

	/*
		Optional volume driver for the container
	*/
	VolumeDriver *string

	/*
		Mount volumes from the specified container(s)
	*/
	VolumesFrom []string

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerContainerCreateCmd is wrapper of 'docker container create'
------------------------------
create [OPTIONS] IMAGE [COMMAND] [ARG...]
Create a new container
------------------------------
*/
func DockerContainerCreateCmd(opt DockerContainerCreateOption, args []string) *exec.Cmd {
	cargs := []string{"container", "create"}
	if opt.AddHost != nil {
		for _, str := range opt.AddHost {
			cargs = append(cargs, "--add-host")
			cargs = append(cargs, str)
		}
	}
	if opt.Attach != nil {
		for _, str := range opt.Attach {
			cargs = append(cargs, "--attach")
			cargs = append(cargs, str)
		}
	}
	if opt.BlkioWeight != nil {
		cargs = append(cargs, "--blkio-weight="+fmt.Sprint(*opt.BlkioWeight))
	}
	if opt.BlkioWeightDevice != nil {
		for _, str := range opt.BlkioWeightDevice {
			cargs = append(cargs, "--blkio-weight-device")
			cargs = append(cargs, str)
		}
	}
	if opt.CapAdd != nil {
		for _, str := range opt.CapAdd {
			cargs = append(cargs, "--cap-add")
			cargs = append(cargs, str)
		}
	}
	if opt.CapDrop != nil {
		for _, str := range opt.CapDrop {
			cargs = append(cargs, "--cap-drop")
			cargs = append(cargs, str)
		}
	}
	if opt.CgroupParent != nil {
		cargs = append(cargs, "--cgroup-parent="+fmt.Sprint(*opt.CgroupParent))
	}
	if opt.Cgroupns != nil {
		cargs = append(cargs, "--cgroupns="+fmt.Sprint(*opt.Cgroupns))
	}
	if opt.Cidfile != nil {
		cargs = append(cargs, "--cidfile="+fmt.Sprint(*opt.Cidfile))
	}
	if opt.CpuCount != nil {
		cargs = append(cargs, "--cpu-count="+fmt.Sprint(*opt.CpuCount))
	}
	if opt.CpuPercent != nil {
		cargs = append(cargs, "--cpu-percent="+fmt.Sprint(*opt.CpuPercent))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuRtPeriod != nil {
		cargs = append(cargs, "--cpu-rt-period="+fmt.Sprint(*opt.CpuRtPeriod))
	}
	if opt.CpuRtRuntime != nil {
		cargs = append(cargs, "--cpu-rt-runtime="+fmt.Sprint(*opt.CpuRtRuntime))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.Cpus != nil {
		cargs = append(cargs, "--cpus="+fmt.Sprint(*opt.Cpus))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.Device != nil {
		for _, str := range opt.Device {
			cargs = append(cargs, "--device")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceCgroupRule != nil {
		for _, str := range opt.DeviceCgroupRule {
			cargs = append(cargs, "--device-cgroup-rule")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadBps != nil {
		for _, str := range opt.DeviceReadBps {
			cargs = append(cargs, "--device-read-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadIops != nil {
		for _, str := range opt.DeviceReadIops {
			cargs = append(cargs, "--device-read-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteBps != nil {
		for _, str := range opt.DeviceWriteBps {
			cargs = append(cargs, "--device-write-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteIops != nil {
		for _, str := range opt.DeviceWriteIops {
			cargs = append(cargs, "--device-write-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Dns != nil {
		for _, str := range opt.Dns {
			cargs = append(cargs, "--dns")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOpt != nil {
		for _, str := range opt.DnsOpt {
			cargs = append(cargs, "--dns-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOption != nil {
		for _, str := range opt.DnsOption {
			cargs = append(cargs, "--dns-option")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsSearch != nil {
		for _, str := range opt.DnsSearch {
			cargs = append(cargs, "--dns-search")
			cargs = append(cargs, str)
		}
	}
	if opt.Domainname != nil {
		cargs = append(cargs, "--domainname="+fmt.Sprint(*opt.Domainname))
	}
	if opt.Entrypoint != nil {
		cargs = append(cargs, "--entrypoint="+fmt.Sprint(*opt.Entrypoint))
	}
	if opt.Env != nil {
		for _, str := range opt.Env {
			cargs = append(cargs, "--env")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvFile != nil {
		for _, str := range opt.EnvFile {
			cargs = append(cargs, "--env-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Expose != nil {
		for _, str := range opt.Expose {
			cargs = append(cargs, "--expose")
			cargs = append(cargs, str)
		}
	}
	if opt.Gpus != nil {
		cargs = append(cargs, "--gpus="+fmt.Sprint(*opt.Gpus))
	}
	if opt.GroupAdd != nil {
		for _, str := range opt.GroupAdd {
			cargs = append(cargs, "--group-add")
			cargs = append(cargs, str)
		}
	}
	if opt.HealthCmd != nil {
		cargs = append(cargs, "--health-cmd="+fmt.Sprint(*opt.HealthCmd))
	}
	if opt.HealthInterval != nil {
		cargs = append(cargs, "--health-interval="+fmt.Sprint(*opt.HealthInterval))
	}
	if opt.HealthRetries != nil {
		cargs = append(cargs, "--health-retries="+fmt.Sprint(*opt.HealthRetries))
	}
	if opt.HealthStartPeriod != nil {
		cargs = append(cargs, "--health-start-period="+fmt.Sprint(*opt.HealthStartPeriod))
	}
	if opt.HealthTimeout != nil {
		cargs = append(cargs, "--health-timeout="+fmt.Sprint(*opt.HealthTimeout))
	}
	if opt.Help != nil {
		cargs = append(cargs, "--help="+fmt.Sprint(*opt.Help))
	}
	if opt.Hostname != nil {
		cargs = append(cargs, "--hostname="+fmt.Sprint(*opt.Hostname))
	}
	if opt.Init != nil {
		cargs = append(cargs, "--init="+fmt.Sprint(*opt.Init))
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	if opt.IoMaxbandwidth != nil {
		cargs = append(cargs, "--io-maxbandwidth="+fmt.Sprint(*opt.IoMaxbandwidth))
	}
	if opt.IoMaxiops != nil {
		cargs = append(cargs, "--io-maxiops="+fmt.Sprint(*opt.IoMaxiops))
	}
	if opt.Ip != nil {
		cargs = append(cargs, "--ip="+fmt.Sprint(*opt.Ip))
	}
	if opt.Ip6 != nil {
		cargs = append(cargs, "--ip6="+fmt.Sprint(*opt.Ip6))
	}
	if opt.Ipc != nil {
		cargs = append(cargs, "--ipc="+fmt.Sprint(*opt.Ipc))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.KernelMemory != nil {
		cargs = append(cargs, "--kernel-memory="+fmt.Sprint(*opt.KernelMemory))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.LabelFile != nil {
		for _, str := range opt.LabelFile {
			cargs = append(cargs, "--label-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Link != nil {
		for _, str := range opt.Link {
			cargs = append(cargs, "--link")
			cargs = append(cargs, str)
		}
	}
	if opt.LinkLocalIp != nil {
		for _, str := range opt.LinkLocalIp {
			cargs = append(cargs, "--link-local-ip")
			cargs = append(cargs, str)
		}
	}
	if opt.LogDriver != nil {
		cargs = append(cargs, "--log-driver="+fmt.Sprint(*opt.LogDriver))
	}
	if opt.LogOpt != nil {
		for _, str := range opt.LogOpt {
			cargs = append(cargs, "--log-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.MacAddress != nil {
		cargs = append(cargs, "--mac-address="+fmt.Sprint(*opt.MacAddress))
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemoryReservation != nil {
		cargs = append(cargs, "--memory-reservation="+fmt.Sprint(*opt.MemoryReservation))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.MemorySwappiness != nil {
		cargs = append(cargs, "--memory-swappiness="+fmt.Sprint(*opt.MemorySwappiness))
	}
	if opt.Mount != nil {
		cargs = append(cargs, "--mount="+fmt.Sprint(*opt.Mount))
	}
	if opt.Name != nil {
		cargs = append(cargs, "--name="+fmt.Sprint(*opt.Name))
	}
	if opt.Net != nil {
		cargs = append(cargs, "--net="+fmt.Sprint(*opt.Net))
	}
	if opt.NetAlias != nil {
		for _, str := range opt.NetAlias {
			cargs = append(cargs, "--net-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NetworkAlias != nil {
		for _, str := range opt.NetworkAlias {
			cargs = append(cargs, "--network-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.NoHealthcheck != nil {
		cargs = append(cargs, "--no-healthcheck="+fmt.Sprint(*opt.NoHealthcheck))
	}
	if opt.OomKillDisable != nil {
		cargs = append(cargs, "--oom-kill-disable="+fmt.Sprint(*opt.OomKillDisable))
	}
	if opt.OomScoreAdj != nil {
		cargs = append(cargs, "--oom-score-adj="+fmt.Sprint(*opt.OomScoreAdj))
	}
	if opt.Pid != nil {
		cargs = append(cargs, "--pid="+fmt.Sprint(*opt.Pid))
	}
	if opt.PidsLimit != nil {
		cargs = append(cargs, "--pids-limit="+fmt.Sprint(*opt.PidsLimit))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Privileged != nil {
		cargs = append(cargs, "--privileged="+fmt.Sprint(*opt.Privileged))
	}
	if opt.Publish != nil {
		for _, str := range opt.Publish {
			cargs = append(cargs, "--publish")
			cargs = append(cargs, str)
		}
	}
	if opt.PublishAll != nil {
		cargs = append(cargs, "--publish-all="+fmt.Sprint(*opt.PublishAll))
	}
	if opt.Pull != nil {
		cargs = append(cargs, "--pull="+fmt.Sprint(*opt.Pull))
	}
	if opt.ReadOnly != nil {
		cargs = append(cargs, "--read-only="+fmt.Sprint(*opt.ReadOnly))
	}
	if opt.Restart != nil {
		cargs = append(cargs, "--restart="+fmt.Sprint(*opt.Restart))
	}
	if opt.Rm != nil {
		cargs = append(cargs, "--rm="+fmt.Sprint(*opt.Rm))
	}
	if opt.Runtime != nil {
		cargs = append(cargs, "--runtime="+fmt.Sprint(*opt.Runtime))
	}
	if opt.SecurityOpt != nil {
		for _, str := range opt.SecurityOpt {
			cargs = append(cargs, "--security-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.ShmSize != nil {
		cargs = append(cargs, "--shm-size="+fmt.Sprint(*opt.ShmSize))
	}
	if opt.StopSignal != nil {
		cargs = append(cargs, "--stop-signal="+fmt.Sprint(*opt.StopSignal))
	}
	if opt.StopTimeout != nil {
		cargs = append(cargs, "--stop-timeout="+fmt.Sprint(*opt.StopTimeout))
	}
	if opt.StorageOpt != nil {
		for _, str := range opt.StorageOpt {
			cargs = append(cargs, "--storage-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.Sysctl != nil {
		for key, val := range opt.Sysctl {
			cargs = append(cargs, "--sysctl")
			cargs = append(cargs, key+"="+val)
		}
	}
	if opt.Tmpfs != nil {
		for _, str := range opt.Tmpfs {
			cargs = append(cargs, "--tmpfs")
			cargs = append(cargs, str)
		}
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.Userns != nil {
		cargs = append(cargs, "--userns="+fmt.Sprint(*opt.Userns))
	}
	if opt.Uts != nil {
		cargs = append(cargs, "--uts="+fmt.Sprint(*opt.Uts))
	}
	if opt.Volume != nil {
		for _, str := range opt.Volume {
			cargs = append(cargs, "--volume")
			cargs = append(cargs, str)
		}
	}
	if opt.VolumeDriver != nil {
		cargs = append(cargs, "--volume-driver="+fmt.Sprint(*opt.VolumeDriver))
	}
	if opt.VolumesFrom != nil {
		for _, str := range opt.VolumesFrom {
			cargs = append(cargs, "--volumes-from")
			cargs = append(cargs, str)
		}
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerDiffCmd is wrapper of 'docker container diff'
------------------------------
diff CONTAINER
Inspect changes to files or directories on a container's filesystem
------------------------------
*/
func DockerContainerDiffCmd(args []string) *exec.Cmd {
	cargs := []string{"container", "diff"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerExecOption struct {
	/*
		Detached mode: run command in the background
	*/
	Detach *bool

	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Set environment variables
	*/
	Env []string

	/*
		Read in a file of environment variables
	*/
	EnvFile []string

	/*
		Keep STDIN open even if not attached
	*/
	Interactive *bool

	/*
		Give extended privileges to the command
	*/
	Privileged *bool

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerContainerExecCmd is wrapper of 'docker container exec'
------------------------------
exec [OPTIONS] CONTAINER COMMAND [ARG...]
Run a command in a running container
------------------------------
*/
func DockerContainerExecCmd(opt DockerContainerExecOption, args []string) *exec.Cmd {
	cargs := []string{"container", "exec"}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.Env != nil {
		for _, str := range opt.Env {
			cargs = append(cargs, "--env")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvFile != nil {
		for _, str := range opt.EnvFile {
			cargs = append(cargs, "--env-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	if opt.Privileged != nil {
		cargs = append(cargs, "--privileged="+fmt.Sprint(*opt.Privileged))
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerExportOption struct {
	/*
		Write to a file, instead of STDOUT
	*/
	Output *string
}

/*
DockerContainerExportCmd is wrapper of 'docker container export'
------------------------------
export [OPTIONS] CONTAINER
Export a container's filesystem as a tar archive
------------------------------
*/
func DockerContainerExportCmd(opt DockerContainerExportOption, args []string) *exec.Cmd {
	cargs := []string{"container", "export"}
	if opt.Output != nil {
		cargs = append(cargs, "--output="+fmt.Sprint(*opt.Output))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Display total file sizes
	*/
	Size *bool
}

/*
DockerContainerInspectCmd is wrapper of 'docker container inspect'
------------------------------
inspect [OPTIONS] CONTAINER [CONTAINER...]
Display detailed information on one or more containers
------------------------------
*/
func DockerContainerInspectCmd(opt DockerContainerInspectOption, args []string) *exec.Cmd {
	cargs := []string{"container", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Size != nil {
		cargs = append(cargs, "--size="+fmt.Sprint(*opt.Size))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerKillOption struct {
	/*
		Signal to send to the container
	*/
	Signal *string
}

/*
DockerContainerKillCmd is wrapper of 'docker container kill'
------------------------------
kill [OPTIONS] CONTAINER [CONTAINER...]
Kill one or more running containers
------------------------------
*/
func DockerContainerKillCmd(opt DockerContainerKillOption, args []string) *exec.Cmd {
	cargs := []string{"container", "kill"}
	if opt.Signal != nil {
		cargs = append(cargs, "--signal="+fmt.Sprint(*opt.Signal))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerLogsOption struct {
	/*
		Show extra details provided to logs
	*/
	Details *bool

	/*
		Follow log output
	*/
	Follow *bool

	/*
		Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
	*/
	Since *string

	/*
		Number of lines to show from the end of the logs
	*/
	Tail *string

	/*
		Show timestamps
	*/
	Timestamps *bool

	/*
		Show logs before a timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
	*/
	Until *string
}

/*
DockerContainerLogsCmd is wrapper of 'docker container logs'
------------------------------
logs [OPTIONS] CONTAINER
Fetch the logs of a container
------------------------------
*/
func DockerContainerLogsCmd(opt DockerContainerLogsOption, args []string) *exec.Cmd {
	cargs := []string{"container", "logs"}
	if opt.Details != nil {
		cargs = append(cargs, "--details="+fmt.Sprint(*opt.Details))
	}
	if opt.Follow != nil {
		cargs = append(cargs, "--follow="+fmt.Sprint(*opt.Follow))
	}
	if opt.Since != nil {
		cargs = append(cargs, "--since="+fmt.Sprint(*opt.Since))
	}
	if opt.Tail != nil {
		cargs = append(cargs, "--tail="+fmt.Sprint(*opt.Tail))
	}
	if opt.Timestamps != nil {
		cargs = append(cargs, "--timestamps="+fmt.Sprint(*opt.Timestamps))
	}
	if opt.Until != nil {
		cargs = append(cargs, "--until="+fmt.Sprint(*opt.Until))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerLsOption struct {
	/*
		Show all containers (default shows just running)
	*/
	All *bool

	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print containers using a Go template
	*/
	Format *string

	/*
		Show n last created containers (includes all states)
	*/
	Last *int

	/*
		Show the latest created container (includes all states)
	*/
	Latest *bool

	/*
		Don't truncate output
	*/
	NoTrunc *bool

	/*
		Only display container IDs
	*/
	Quiet *bool

	/*
		Display total file sizes
	*/
	Size *bool
}

/*
DockerContainerLsCmd is wrapper of 'docker container ls'
------------------------------
ls [OPTIONS]
List containers
------------------------------
*/
func DockerContainerLsCmd(opt DockerContainerLsOption, args []string) *exec.Cmd {
	cargs := []string{"container", "ls"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Last != nil {
		cargs = append(cargs, "--last="+fmt.Sprint(*opt.Last))
	}
	if opt.Latest != nil {
		cargs = append(cargs, "--latest="+fmt.Sprint(*opt.Latest))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Size != nil {
		cargs = append(cargs, "--size="+fmt.Sprint(*opt.Size))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerPauseCmd is wrapper of 'docker container pause'
------------------------------
pause CONTAINER [CONTAINER...]
Pause all processes within one or more containers
------------------------------
*/
func DockerContainerPauseCmd(args []string) *exec.Cmd {
	cargs := []string{"container", "pause"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerPortCmd is wrapper of 'docker container port'
------------------------------
port CONTAINER [PRIVATE_PORT[/PROTO]]
List port mappings or a specific mapping for the container
------------------------------
*/
func DockerContainerPortCmd(args []string) *exec.Cmd {
	cargs := []string{"container", "port"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerPruneOption struct {
	/*
		Provide filter values (e.g. 'until=<timestamp>')
	*/
	Filter *string

	/*
		Do not prompt for confirmation
	*/
	Force *bool
}

/*
DockerContainerPruneCmd is wrapper of 'docker container prune'
------------------------------
prune [OPTIONS]
Remove all stopped containers
------------------------------
*/
func DockerContainerPruneCmd(opt DockerContainerPruneOption, args []string) *exec.Cmd {
	cargs := []string{"container", "prune"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerRenameCmd is wrapper of 'docker container rename'
------------------------------
rename CONTAINER NEW_NAME
Rename a container
------------------------------
*/
func DockerContainerRenameCmd(args []string) *exec.Cmd {
	cargs := []string{"container", "rename"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerRestartOption struct {
	/*
		Seconds to wait for stop before killing the container
	*/
	Time *int
}

/*
DockerContainerRestartCmd is wrapper of 'docker container restart'
------------------------------
restart [OPTIONS] CONTAINER [CONTAINER...]
Restart one or more containers
------------------------------
*/
func DockerContainerRestartCmd(opt DockerContainerRestartOption, args []string) *exec.Cmd {
	cargs := []string{"container", "restart"}
	if opt.Time != nil {
		cargs = append(cargs, "--time="+fmt.Sprint(*opt.Time))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerRmOption struct {
	/*
		Force the removal of a running container (uses SIGKILL)
	*/
	Force *bool

	/*
		Remove the specified link
	*/
	Link *bool

	/*
		Remove anonymous volumes associated with the container
	*/
	Volumes *bool
}

/*
DockerContainerRmCmd is wrapper of 'docker container rm'
------------------------------
rm [OPTIONS] CONTAINER [CONTAINER...]
Remove one or more containers
------------------------------
*/
func DockerContainerRmCmd(opt DockerContainerRmOption, args []string) *exec.Cmd {
	cargs := []string{"container", "rm"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	if opt.Link != nil {
		cargs = append(cargs, "--link="+fmt.Sprint(*opt.Link))
	}
	if opt.Volumes != nil {
		cargs = append(cargs, "--volumes="+fmt.Sprint(*opt.Volumes))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerRunOption struct {
	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	AddHost []string

	/*
		Attach to STDIN, STDOUT or STDERR
	*/
	Attach []string

	/*
		Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
	*/
	BlkioWeight *uint16

	/*
		Block IO weight (relative device weight)
	*/
	BlkioWeightDevice []string

	/*
		Add Linux capabilities
	*/
	CapAdd []string

	/*
		Drop Linux capabilities
	*/
	CapDrop []string

	/*
		Optional parent cgroup for the container
	*/
	CgroupParent *string

	/*
			Cgroup namespace to use (host|private)
		'host':    Run the container in the Docker host's cgroup namespace
		'private': Run the container in its own private cgroup namespace
		'':        Use the cgroup namespace as configured by the
		           default-cgroupns-mode option on the daemon (default)
	*/
	Cgroupns *string

	/*
		Write the container ID to the file
	*/
	Cidfile *string

	/*
		CPU count (Windows only)
	*/
	CpuCount *int64

	/*
		CPU percent (Windows only)
	*/
	CpuPercent *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		Limit CPU real-time period in microseconds
	*/
	CpuRtPeriod *int64

	/*
		Limit CPU real-time runtime in microseconds
	*/
	CpuRtRuntime *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		Number of CPUs
	*/
	Cpus *string

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Run container in background and print container ID
	*/
	Detach *bool

	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Add a host device to the container
	*/
	Device []string

	/*
		Add a rule to the cgroup allowed devices list
	*/
	DeviceCgroupRule []string

	/*
		Limit read rate (bytes per second) from a device
	*/
	DeviceReadBps []string

	/*
		Limit read rate (IO per second) from a device
	*/
	DeviceReadIops []string

	/*
		Limit write rate (bytes per second) to a device
	*/
	DeviceWriteBps []string

	/*
		Limit write rate (IO per second) to a device
	*/
	DeviceWriteIops []string

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Set custom DNS servers
	*/
	Dns []string

	/*
		Set DNS options
	*/
	DnsOpt []string

	/*
		Set DNS options
	*/
	DnsOption []string

	/*
		Set custom DNS search domains
	*/
	DnsSearch []string

	/*
		Container NIS domain name
	*/
	Domainname *string

	/*
		Overwrite the default ENTRYPOINT of the image
	*/
	Entrypoint *string

	/*
		Set environment variables
	*/
	Env []string

	/*
		Read in a file of environment variables
	*/
	EnvFile []string

	/*
		Expose a port or a range of ports
	*/
	Expose []string

	/*
		GPU devices to add to the container ('all' to pass all GPUs)
	*/
	Gpus *string

	/*
		Add additional groups to join
	*/
	GroupAdd []string

	/*
		Command to run to check health
	*/
	HealthCmd *string

	/*
		Time between running the check (ms|s|m|h) (default 0s)
	*/
	HealthInterval *string

	/*
		Consecutive failures needed to report unhealthy
	*/
	HealthRetries *int

	/*
		Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
	*/
	HealthStartPeriod *string

	/*
		Maximum time to allow one check to run (ms|s|m|h) (default 0s)
	*/
	HealthTimeout *string

	/*
		Print usage
	*/
	Help *bool

	/*
		Container host name
	*/
	Hostname *string

	/*
		Run an init inside the container that forwards signals and reaps processes
	*/
	Init *bool

	/*
		Keep STDIN open even if not attached
	*/
	Interactive *bool

	/*
		Maximum IO bandwidth limit for the system drive (Windows only)
	*/
	IoMaxbandwidth *string

	/*
		Maximum IOps limit for the system drive (Windows only)
	*/
	IoMaxiops *uint64

	/*
		IPv4 address (e.g., 172.30.100.104)
	*/
	Ip *string

	/*
		IPv6 address (e.g., 2001:db8::33)
	*/
	Ip6 *string

	/*
		IPC mode to use
	*/
	Ipc *string

	/*
		Container isolation technology
	*/
	Isolation *string

	/*
		Kernel memory limit
	*/
	KernelMemory *string

	/*
		Set meta data on a container
	*/
	Label []string

	/*
		Read in a line delimited file of labels
	*/
	LabelFile []string

	/*
		Add link to another container
	*/
	Link []string

	/*
		Container IPv4/IPv6 link-local addresses
	*/
	LinkLocalIp []string

	/*
		Logging driver for the container
	*/
	LogDriver *string

	/*
		Log driver options
	*/
	LogOpt []string

	/*
		Container MAC address (e.g., 92:d0:c6:0a:29:33)
	*/
	MacAddress *string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Memory soft limit
	*/
	MemoryReservation *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Tune container memory swappiness (0 to 100)
	*/
	MemorySwappiness *int64

	/*
		Attach a filesystem mount to the container
	*/
	Mount *string

	/*
		Assign a name to the container
	*/
	Name *string

	/*
		Connect a container to a network
	*/
	Net *string

	/*
		Add network-scoped alias for the container
	*/
	NetAlias []string

	/*
		Connect a container to a network
	*/
	Network *string

	/*
		Add network-scoped alias for the container
	*/
	NetworkAlias []string

	/*
		Disable any container-specified HEALTHCHECK
	*/
	NoHealthcheck *bool

	/*
		Disable OOM Killer
	*/
	OomKillDisable *bool

	/*
		Tune host's OOM preferences (-1000 to 1000)
	*/
	OomScoreAdj *int

	/*
		PID namespace to use
	*/
	Pid *string

	/*
		Tune container pids limit (set -1 for unlimited)
	*/
	PidsLimit *int64

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Give extended privileges to this container
	*/
	Privileged *bool

	/*
		Publish a container's port(s) to the host
	*/
	Publish []string

	/*
		Publish all exposed ports to random ports
	*/
	PublishAll *bool

	/*
		Pull image before running ("always"|"missing"|"never")
	*/
	Pull *string

	/*
		Mount the container's root filesystem as read only
	*/
	ReadOnly *bool

	/*
		Restart policy to apply when a container exits
	*/
	Restart *string

	/*
		Automatically remove the container when it exits
	*/
	Rm *bool

	/*
		Runtime to use for this container
	*/
	Runtime *string

	/*
		Security Options
	*/
	SecurityOpt []string

	/*
		Size of /dev/shm
	*/
	ShmSize *string

	/*
		Proxy received signals to the process
	*/
	SigProxy *bool

	/*
		Signal to stop a container
	*/
	StopSignal *string

	/*
		Timeout (in seconds) to stop a container
	*/
	StopTimeout *int

	/*
		Storage driver options for the container
	*/
	StorageOpt []string

	/*
		Sysctl options
	*/
	Sysctl map[string]string

	/*
		Mount a tmpfs directory
	*/
	Tmpfs []string

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Ulimit options
	*/
	Ulimit *string

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		User namespace to use
	*/
	Userns *string

	/*
		UTS namespace to use
	*/
	Uts *string

	/*
		Bind mount a volume
	*/
	Volume []string

	/*
		Optional volume driver for the container
	*/
	VolumeDriver *string

	/*
		Mount volumes from the specified container(s)
	*/
	VolumesFrom []string

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerContainerRunCmd is wrapper of 'docker container run'
------------------------------
run [OPTIONS] IMAGE [COMMAND] [ARG...]
Run a command in a new container
------------------------------
*/
func DockerContainerRunCmd(opt DockerContainerRunOption, args []string) *exec.Cmd {
	cargs := []string{"container", "run"}
	if opt.AddHost != nil {
		for _, str := range opt.AddHost {
			cargs = append(cargs, "--add-host")
			cargs = append(cargs, str)
		}
	}
	if opt.Attach != nil {
		for _, str := range opt.Attach {
			cargs = append(cargs, "--attach")
			cargs = append(cargs, str)
		}
	}
	if opt.BlkioWeight != nil {
		cargs = append(cargs, "--blkio-weight="+fmt.Sprint(*opt.BlkioWeight))
	}
	if opt.BlkioWeightDevice != nil {
		for _, str := range opt.BlkioWeightDevice {
			cargs = append(cargs, "--blkio-weight-device")
			cargs = append(cargs, str)
		}
	}
	if opt.CapAdd != nil {
		for _, str := range opt.CapAdd {
			cargs = append(cargs, "--cap-add")
			cargs = append(cargs, str)
		}
	}
	if opt.CapDrop != nil {
		for _, str := range opt.CapDrop {
			cargs = append(cargs, "--cap-drop")
			cargs = append(cargs, str)
		}
	}
	if opt.CgroupParent != nil {
		cargs = append(cargs, "--cgroup-parent="+fmt.Sprint(*opt.CgroupParent))
	}
	if opt.Cgroupns != nil {
		cargs = append(cargs, "--cgroupns="+fmt.Sprint(*opt.Cgroupns))
	}
	if opt.Cidfile != nil {
		cargs = append(cargs, "--cidfile="+fmt.Sprint(*opt.Cidfile))
	}
	if opt.CpuCount != nil {
		cargs = append(cargs, "--cpu-count="+fmt.Sprint(*opt.CpuCount))
	}
	if opt.CpuPercent != nil {
		cargs = append(cargs, "--cpu-percent="+fmt.Sprint(*opt.CpuPercent))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuRtPeriod != nil {
		cargs = append(cargs, "--cpu-rt-period="+fmt.Sprint(*opt.CpuRtPeriod))
	}
	if opt.CpuRtRuntime != nil {
		cargs = append(cargs, "--cpu-rt-runtime="+fmt.Sprint(*opt.CpuRtRuntime))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.Cpus != nil {
		cargs = append(cargs, "--cpus="+fmt.Sprint(*opt.Cpus))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.Device != nil {
		for _, str := range opt.Device {
			cargs = append(cargs, "--device")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceCgroupRule != nil {
		for _, str := range opt.DeviceCgroupRule {
			cargs = append(cargs, "--device-cgroup-rule")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadBps != nil {
		for _, str := range opt.DeviceReadBps {
			cargs = append(cargs, "--device-read-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadIops != nil {
		for _, str := range opt.DeviceReadIops {
			cargs = append(cargs, "--device-read-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteBps != nil {
		for _, str := range opt.DeviceWriteBps {
			cargs = append(cargs, "--device-write-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteIops != nil {
		for _, str := range opt.DeviceWriteIops {
			cargs = append(cargs, "--device-write-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Dns != nil {
		for _, str := range opt.Dns {
			cargs = append(cargs, "--dns")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOpt != nil {
		for _, str := range opt.DnsOpt {
			cargs = append(cargs, "--dns-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOption != nil {
		for _, str := range opt.DnsOption {
			cargs = append(cargs, "--dns-option")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsSearch != nil {
		for _, str := range opt.DnsSearch {
			cargs = append(cargs, "--dns-search")
			cargs = append(cargs, str)
		}
	}
	if opt.Domainname != nil {
		cargs = append(cargs, "--domainname="+fmt.Sprint(*opt.Domainname))
	}
	if opt.Entrypoint != nil {
		cargs = append(cargs, "--entrypoint="+fmt.Sprint(*opt.Entrypoint))
	}
	if opt.Env != nil {
		for _, str := range opt.Env {
			cargs = append(cargs, "--env")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvFile != nil {
		for _, str := range opt.EnvFile {
			cargs = append(cargs, "--env-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Expose != nil {
		for _, str := range opt.Expose {
			cargs = append(cargs, "--expose")
			cargs = append(cargs, str)
		}
	}
	if opt.Gpus != nil {
		cargs = append(cargs, "--gpus="+fmt.Sprint(*opt.Gpus))
	}
	if opt.GroupAdd != nil {
		for _, str := range opt.GroupAdd {
			cargs = append(cargs, "--group-add")
			cargs = append(cargs, str)
		}
	}
	if opt.HealthCmd != nil {
		cargs = append(cargs, "--health-cmd="+fmt.Sprint(*opt.HealthCmd))
	}
	if opt.HealthInterval != nil {
		cargs = append(cargs, "--health-interval="+fmt.Sprint(*opt.HealthInterval))
	}
	if opt.HealthRetries != nil {
		cargs = append(cargs, "--health-retries="+fmt.Sprint(*opt.HealthRetries))
	}
	if opt.HealthStartPeriod != nil {
		cargs = append(cargs, "--health-start-period="+fmt.Sprint(*opt.HealthStartPeriod))
	}
	if opt.HealthTimeout != nil {
		cargs = append(cargs, "--health-timeout="+fmt.Sprint(*opt.HealthTimeout))
	}
	if opt.Help != nil {
		cargs = append(cargs, "--help="+fmt.Sprint(*opt.Help))
	}
	if opt.Hostname != nil {
		cargs = append(cargs, "--hostname="+fmt.Sprint(*opt.Hostname))
	}
	if opt.Init != nil {
		cargs = append(cargs, "--init="+fmt.Sprint(*opt.Init))
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	if opt.IoMaxbandwidth != nil {
		cargs = append(cargs, "--io-maxbandwidth="+fmt.Sprint(*opt.IoMaxbandwidth))
	}
	if opt.IoMaxiops != nil {
		cargs = append(cargs, "--io-maxiops="+fmt.Sprint(*opt.IoMaxiops))
	}
	if opt.Ip != nil {
		cargs = append(cargs, "--ip="+fmt.Sprint(*opt.Ip))
	}
	if opt.Ip6 != nil {
		cargs = append(cargs, "--ip6="+fmt.Sprint(*opt.Ip6))
	}
	if opt.Ipc != nil {
		cargs = append(cargs, "--ipc="+fmt.Sprint(*opt.Ipc))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.KernelMemory != nil {
		cargs = append(cargs, "--kernel-memory="+fmt.Sprint(*opt.KernelMemory))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.LabelFile != nil {
		for _, str := range opt.LabelFile {
			cargs = append(cargs, "--label-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Link != nil {
		for _, str := range opt.Link {
			cargs = append(cargs, "--link")
			cargs = append(cargs, str)
		}
	}
	if opt.LinkLocalIp != nil {
		for _, str := range opt.LinkLocalIp {
			cargs = append(cargs, "--link-local-ip")
			cargs = append(cargs, str)
		}
	}
	if opt.LogDriver != nil {
		cargs = append(cargs, "--log-driver="+fmt.Sprint(*opt.LogDriver))
	}
	if opt.LogOpt != nil {
		for _, str := range opt.LogOpt {
			cargs = append(cargs, "--log-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.MacAddress != nil {
		cargs = append(cargs, "--mac-address="+fmt.Sprint(*opt.MacAddress))
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemoryReservation != nil {
		cargs = append(cargs, "--memory-reservation="+fmt.Sprint(*opt.MemoryReservation))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.MemorySwappiness != nil {
		cargs = append(cargs, "--memory-swappiness="+fmt.Sprint(*opt.MemorySwappiness))
	}
	if opt.Mount != nil {
		cargs = append(cargs, "--mount="+fmt.Sprint(*opt.Mount))
	}
	if opt.Name != nil {
		cargs = append(cargs, "--name="+fmt.Sprint(*opt.Name))
	}
	if opt.Net != nil {
		cargs = append(cargs, "--net="+fmt.Sprint(*opt.Net))
	}
	if opt.NetAlias != nil {
		for _, str := range opt.NetAlias {
			cargs = append(cargs, "--net-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NetworkAlias != nil {
		for _, str := range opt.NetworkAlias {
			cargs = append(cargs, "--network-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.NoHealthcheck != nil {
		cargs = append(cargs, "--no-healthcheck="+fmt.Sprint(*opt.NoHealthcheck))
	}
	if opt.OomKillDisable != nil {
		cargs = append(cargs, "--oom-kill-disable="+fmt.Sprint(*opt.OomKillDisable))
	}
	if opt.OomScoreAdj != nil {
		cargs = append(cargs, "--oom-score-adj="+fmt.Sprint(*opt.OomScoreAdj))
	}
	if opt.Pid != nil {
		cargs = append(cargs, "--pid="+fmt.Sprint(*opt.Pid))
	}
	if opt.PidsLimit != nil {
		cargs = append(cargs, "--pids-limit="+fmt.Sprint(*opt.PidsLimit))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Privileged != nil {
		cargs = append(cargs, "--privileged="+fmt.Sprint(*opt.Privileged))
	}
	if opt.Publish != nil {
		for _, str := range opt.Publish {
			cargs = append(cargs, "--publish")
			cargs = append(cargs, str)
		}
	}
	if opt.PublishAll != nil {
		cargs = append(cargs, "--publish-all="+fmt.Sprint(*opt.PublishAll))
	}
	if opt.Pull != nil {
		cargs = append(cargs, "--pull="+fmt.Sprint(*opt.Pull))
	}
	if opt.ReadOnly != nil {
		cargs = append(cargs, "--read-only="+fmt.Sprint(*opt.ReadOnly))
	}
	if opt.Restart != nil {
		cargs = append(cargs, "--restart="+fmt.Sprint(*opt.Restart))
	}
	if opt.Rm != nil {
		cargs = append(cargs, "--rm="+fmt.Sprint(*opt.Rm))
	}
	if opt.Runtime != nil {
		cargs = append(cargs, "--runtime="+fmt.Sprint(*opt.Runtime))
	}
	if opt.SecurityOpt != nil {
		for _, str := range opt.SecurityOpt {
			cargs = append(cargs, "--security-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.ShmSize != nil {
		cargs = append(cargs, "--shm-size="+fmt.Sprint(*opt.ShmSize))
	}
	if opt.SigProxy != nil {
		cargs = append(cargs, "--sig-proxy="+fmt.Sprint(*opt.SigProxy))
	}
	if opt.StopSignal != nil {
		cargs = append(cargs, "--stop-signal="+fmt.Sprint(*opt.StopSignal))
	}
	if opt.StopTimeout != nil {
		cargs = append(cargs, "--stop-timeout="+fmt.Sprint(*opt.StopTimeout))
	}
	if opt.StorageOpt != nil {
		for _, str := range opt.StorageOpt {
			cargs = append(cargs, "--storage-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.Sysctl != nil {
		for key, val := range opt.Sysctl {
			cargs = append(cargs, "--sysctl")
			cargs = append(cargs, key+"="+val)
		}
	}
	if opt.Tmpfs != nil {
		for _, str := range opt.Tmpfs {
			cargs = append(cargs, "--tmpfs")
			cargs = append(cargs, str)
		}
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.Userns != nil {
		cargs = append(cargs, "--userns="+fmt.Sprint(*opt.Userns))
	}
	if opt.Uts != nil {
		cargs = append(cargs, "--uts="+fmt.Sprint(*opt.Uts))
	}
	if opt.Volume != nil {
		for _, str := range opt.Volume {
			cargs = append(cargs, "--volume")
			cargs = append(cargs, str)
		}
	}
	if opt.VolumeDriver != nil {
		cargs = append(cargs, "--volume-driver="+fmt.Sprint(*opt.VolumeDriver))
	}
	if opt.VolumesFrom != nil {
		for _, str := range opt.VolumesFrom {
			cargs = append(cargs, "--volumes-from")
			cargs = append(cargs, str)
		}
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerStartOption struct {
	/*
		Attach STDOUT/STDERR and forward signals
	*/
	Attach *bool

	/*
		Restore from this checkpoint
	*/
	Checkpoint *string

	/*
		Use a custom checkpoint storage directory
	*/
	CheckpointDir *string

	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Attach container's STDIN
	*/
	Interactive *bool
}

/*
DockerContainerStartCmd is wrapper of 'docker container start'
------------------------------
start [OPTIONS] CONTAINER [CONTAINER...]
Start one or more stopped containers
------------------------------
*/
func DockerContainerStartCmd(opt DockerContainerStartOption, args []string) *exec.Cmd {
	cargs := []string{"container", "start"}
	if opt.Attach != nil {
		cargs = append(cargs, "--attach="+fmt.Sprint(*opt.Attach))
	}
	if opt.Checkpoint != nil {
		cargs = append(cargs, "--checkpoint="+fmt.Sprint(*opt.Checkpoint))
	}
	if opt.CheckpointDir != nil {
		cargs = append(cargs, "--checkpoint-dir="+fmt.Sprint(*opt.CheckpointDir))
	}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerStatsOption struct {
	/*
		Show all containers (default shows just running)
	*/
	All *bool

	/*
		Pretty-print images using a Go template
	*/
	Format *string

	/*
		Disable streaming stats and only pull the first result
	*/
	NoStream *bool

	/*
		Do not truncate output
	*/
	NoTrunc *bool
}

/*
DockerContainerStatsCmd is wrapper of 'docker container stats'
------------------------------
stats [OPTIONS] [CONTAINER...]
Display a live stream of container(s) resource usage statistics
------------------------------
*/
func DockerContainerStatsCmd(opt DockerContainerStatsOption, args []string) *exec.Cmd {
	cargs := []string{"container", "stats"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoStream != nil {
		cargs = append(cargs, "--no-stream="+fmt.Sprint(*opt.NoStream))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerStopOption struct {
	/*
		Seconds to wait for stop before killing it
	*/
	Time *int
}

/*
DockerContainerStopCmd is wrapper of 'docker container stop'
------------------------------
stop [OPTIONS] CONTAINER [CONTAINER...]
Stop one or more running containers
------------------------------
*/
func DockerContainerStopCmd(opt DockerContainerStopOption, args []string) *exec.Cmd {
	cargs := []string{"container", "stop"}
	if opt.Time != nil {
		cargs = append(cargs, "--time="+fmt.Sprint(*opt.Time))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerTopCmd is wrapper of 'docker container top'
------------------------------
top CONTAINER [ps OPTIONS]
Display the running processes of a container
------------------------------
*/
func DockerContainerTopCmd(args []string) *exec.Cmd {
	cargs := []string{"container", "top"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerUnpauseCmd is wrapper of 'docker container unpause'
------------------------------
unpause CONTAINER [CONTAINER...]
Unpause all processes within one or more containers
------------------------------
*/
func DockerContainerUnpauseCmd(args []string) *exec.Cmd {
	cargs := []string{"container", "unpause"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContainerUpdateOption struct {
	/*
		Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
	*/
	BlkioWeight *uint16

	/*
		Limit CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		Limit the CPU real-time period in microseconds
	*/
	CpuRtPeriod *int64

	/*
		Limit the CPU real-time runtime in microseconds
	*/
	CpuRtRuntime *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		Number of CPUs
	*/
	Cpus *string

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Kernel memory limit
	*/
	KernelMemory *string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Memory soft limit
	*/
	MemoryReservation *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Tune container pids limit (set -1 for unlimited)
	*/
	PidsLimit *int64

	/*
		Restart policy to apply when a container exits
	*/
	Restart *string
}

/*
DockerContainerUpdateCmd is wrapper of 'docker container update'
------------------------------
update [OPTIONS] CONTAINER [CONTAINER...]
Update configuration of one or more containers
------------------------------
*/
func DockerContainerUpdateCmd(opt DockerContainerUpdateOption, args []string) *exec.Cmd {
	cargs := []string{"container", "update"}
	if opt.BlkioWeight != nil {
		cargs = append(cargs, "--blkio-weight="+fmt.Sprint(*opt.BlkioWeight))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuRtPeriod != nil {
		cargs = append(cargs, "--cpu-rt-period="+fmt.Sprint(*opt.CpuRtPeriod))
	}
	if opt.CpuRtRuntime != nil {
		cargs = append(cargs, "--cpu-rt-runtime="+fmt.Sprint(*opt.CpuRtRuntime))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.Cpus != nil {
		cargs = append(cargs, "--cpus="+fmt.Sprint(*opt.Cpus))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.KernelMemory != nil {
		cargs = append(cargs, "--kernel-memory="+fmt.Sprint(*opt.KernelMemory))
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemoryReservation != nil {
		cargs = append(cargs, "--memory-reservation="+fmt.Sprint(*opt.MemoryReservation))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.PidsLimit != nil {
		cargs = append(cargs, "--pids-limit="+fmt.Sprint(*opt.PidsLimit))
	}
	if opt.Restart != nil {
		cargs = append(cargs, "--restart="+fmt.Sprint(*opt.Restart))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContainerWaitCmd is wrapper of 'docker container wait'
------------------------------
wait CONTAINER [CONTAINER...]
Block until one or more containers stop, then print their exit codes
------------------------------
*/
func DockerContainerWaitCmd(args []string) *exec.Cmd {
	cargs := []string{"container", "wait"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContextCmd is wrapper of 'docker context'
------------------------------
context
Manage contexts
------------------------------
*/
func DockerContextCmd(args []string) *exec.Cmd {
	cargs := []string{"context"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContextCreateOption struct {
	/*
		Default orchestrator for stack operations to use with this context (swarm|kubernetes|all)
	*/
	DefaultStackOrchestrator *string

	/*
		Description of the context
	*/
	Description *string

	/*
		set the docker endpoint
	*/
	Docker *string

	/*
		create context from a named context
	*/
	From *string

	/*
		set the kubernetes endpoint
	*/
	Kubernetes *string
}

/*
DockerContextCreateCmd is wrapper of 'docker context create'
------------------------------
create [OPTIONS] CONTEXT
Create a context
------------------------------
*/
func DockerContextCreateCmd(opt DockerContextCreateOption, args []string) *exec.Cmd {
	cargs := []string{"context", "create"}
	if opt.DefaultStackOrchestrator != nil {
		cargs = append(cargs, "--default-stack-orchestrator="+fmt.Sprint(*opt.DefaultStackOrchestrator))
	}
	if opt.Description != nil {
		cargs = append(cargs, "--description="+fmt.Sprint(*opt.Description))
	}
	if opt.Docker != nil {
		cargs = append(cargs, "--docker="+fmt.Sprint(*opt.Docker))
	}
	if opt.From != nil {
		cargs = append(cargs, "--from="+fmt.Sprint(*opt.From))
	}
	if opt.Kubernetes != nil {
		cargs = append(cargs, "--kubernetes="+fmt.Sprint(*opt.Kubernetes))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContextExportOption struct {
	/*
		Export as a kubeconfig file
	*/
	Kubeconfig *bool
}

/*
DockerContextExportCmd is wrapper of 'docker context export'
------------------------------
export [OPTIONS] CONTEXT [FILE|-]
Export a context to a tar or kubeconfig file
------------------------------
*/
func DockerContextExportCmd(opt DockerContextExportOption, args []string) *exec.Cmd {
	cargs := []string{"context", "export"}
	if opt.Kubeconfig != nil {
		cargs = append(cargs, "--kubeconfig="+fmt.Sprint(*opt.Kubeconfig))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContextImportCmd is wrapper of 'docker context import'
------------------------------
import CONTEXT FILE|-
Import a context from a tar or zip file
------------------------------
*/
func DockerContextImportCmd(args []string) *exec.Cmd {
	cargs := []string{"context", "import"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContextInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string
}

/*
DockerContextInspectCmd is wrapper of 'docker context inspect'
------------------------------
inspect [OPTIONS] [CONTEXT] [CONTEXT...]
Display detailed information on one or more contexts
------------------------------
*/
func DockerContextInspectCmd(opt DockerContextInspectOption, args []string) *exec.Cmd {
	cargs := []string{"context", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContextLsOption struct {
	/*
		Pretty-print contexts using a Go template
	*/
	Format *string

	/*
		Only show context names
	*/
	Quiet *bool
}

/*
DockerContextLsCmd is wrapper of 'docker context ls'
------------------------------
ls [OPTIONS]
List contexts
------------------------------
*/
func DockerContextLsCmd(opt DockerContextLsOption, args []string) *exec.Cmd {
	cargs := []string{"context", "ls"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContextRmOption struct {
	/*
		Force the removal of a context in use
	*/
	Force *bool
}

/*
DockerContextRmCmd is wrapper of 'docker context rm'
------------------------------
rm CONTEXT [CONTEXT...]
Remove one or more contexts
------------------------------
*/
func DockerContextRmCmd(opt DockerContextRmOption, args []string) *exec.Cmd {
	cargs := []string{"context", "rm"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerContextUpdateOption struct {
	/*
		Default orchestrator for stack operations to use with this context (swarm|kubernetes|all)
	*/
	DefaultStackOrchestrator *string

	/*
		Description of the context
	*/
	Description *string

	/*
		set the docker endpoint
	*/
	Docker *string

	/*
		set the kubernetes endpoint
	*/
	Kubernetes *string
}

/*
DockerContextUpdateCmd is wrapper of 'docker context update'
------------------------------
update [OPTIONS] CONTEXT
Update a context
------------------------------
*/
func DockerContextUpdateCmd(opt DockerContextUpdateOption, args []string) *exec.Cmd {
	cargs := []string{"context", "update"}
	if opt.DefaultStackOrchestrator != nil {
		cargs = append(cargs, "--default-stack-orchestrator="+fmt.Sprint(*opt.DefaultStackOrchestrator))
	}
	if opt.Description != nil {
		cargs = append(cargs, "--description="+fmt.Sprint(*opt.Description))
	}
	if opt.Docker != nil {
		cargs = append(cargs, "--docker="+fmt.Sprint(*opt.Docker))
	}
	if opt.Kubernetes != nil {
		cargs = append(cargs, "--kubernetes="+fmt.Sprint(*opt.Kubernetes))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerContextUseCmd is wrapper of 'docker context use'
------------------------------
use CONTEXT
Set the current docker context
------------------------------
*/
func DockerContextUseCmd(args []string) *exec.Cmd {
	cargs := []string{"context", "use"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerCpOption struct {
	/*
		Archive mode (copy all uid/gid information)
	*/
	Archive *bool

	/*
		Always follow symbol link in SRC_PATH
	*/
	FollowLink *bool
}

/*
DockerCpCmd is wrapper of 'docker cp'
------------------------------
cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-

	docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH

Copy files/folders between a container and the local filesystem
------------------------------
*/
func DockerCpCmd(opt DockerCpOption, args []string) *exec.Cmd {
	cargs := []string{"cp"}
	if opt.Archive != nil {
		cargs = append(cargs, "--archive="+fmt.Sprint(*opt.Archive))
	}
	if opt.FollowLink != nil {
		cargs = append(cargs, "--follow-link="+fmt.Sprint(*opt.FollowLink))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerCreateOption struct {
	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	AddHost []string

	/*
		Attach to STDIN, STDOUT or STDERR
	*/
	Attach []string

	/*
		Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
	*/
	BlkioWeight *uint16

	/*
		Block IO weight (relative device weight)
	*/
	BlkioWeightDevice []string

	/*
		Add Linux capabilities
	*/
	CapAdd []string

	/*
		Drop Linux capabilities
	*/
	CapDrop []string

	/*
		Optional parent cgroup for the container
	*/
	CgroupParent *string

	/*
			Cgroup namespace to use (host|private)
		'host':    Run the container in the Docker host's cgroup namespace
		'private': Run the container in its own private cgroup namespace
		'':        Use the cgroup namespace as configured by the
		           default-cgroupns-mode option on the daemon (default)
	*/
	Cgroupns *string

	/*
		Write the container ID to the file
	*/
	Cidfile *string

	/*
		CPU count (Windows only)
	*/
	CpuCount *int64

	/*
		CPU percent (Windows only)
	*/
	CpuPercent *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		Limit CPU real-time period in microseconds
	*/
	CpuRtPeriod *int64

	/*
		Limit CPU real-time runtime in microseconds
	*/
	CpuRtRuntime *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		Number of CPUs
	*/
	Cpus *string

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Add a host device to the container
	*/
	Device []string

	/*
		Add a rule to the cgroup allowed devices list
	*/
	DeviceCgroupRule []string

	/*
		Limit read rate (bytes per second) from a device
	*/
	DeviceReadBps []string

	/*
		Limit read rate (IO per second) from a device
	*/
	DeviceReadIops []string

	/*
		Limit write rate (bytes per second) to a device
	*/
	DeviceWriteBps []string

	/*
		Limit write rate (IO per second) to a device
	*/
	DeviceWriteIops []string

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Set custom DNS servers
	*/
	Dns []string

	/*
		Set DNS options
	*/
	DnsOpt []string

	/*
		Set DNS options
	*/
	DnsOption []string

	/*
		Set custom DNS search domains
	*/
	DnsSearch []string

	/*
		Container NIS domain name
	*/
	Domainname *string

	/*
		Overwrite the default ENTRYPOINT of the image
	*/
	Entrypoint *string

	/*
		Set environment variables
	*/
	Env []string

	/*
		Read in a file of environment variables
	*/
	EnvFile []string

	/*
		Expose a port or a range of ports
	*/
	Expose []string

	/*
		GPU devices to add to the container ('all' to pass all GPUs)
	*/
	Gpus *string

	/*
		Add additional groups to join
	*/
	GroupAdd []string

	/*
		Command to run to check health
	*/
	HealthCmd *string

	/*
		Time between running the check (ms|s|m|h) (default 0s)
	*/
	HealthInterval *string

	/*
		Consecutive failures needed to report unhealthy
	*/
	HealthRetries *int

	/*
		Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
	*/
	HealthStartPeriod *string

	/*
		Maximum time to allow one check to run (ms|s|m|h) (default 0s)
	*/
	HealthTimeout *string

	/*
		Print usage
	*/
	Help *bool

	/*
		Container host name
	*/
	Hostname *string

	/*
		Run an init inside the container that forwards signals and reaps processes
	*/
	Init *bool

	/*
		Keep STDIN open even if not attached
	*/
	Interactive *bool

	/*
		Maximum IO bandwidth limit for the system drive (Windows only)
	*/
	IoMaxbandwidth *string

	/*
		Maximum IOps limit for the system drive (Windows only)
	*/
	IoMaxiops *uint64

	/*
		IPv4 address (e.g., 172.30.100.104)
	*/
	Ip *string

	/*
		IPv6 address (e.g., 2001:db8::33)
	*/
	Ip6 *string

	/*
		IPC mode to use
	*/
	Ipc *string

	/*
		Container isolation technology
	*/
	Isolation *string

	/*
		Kernel memory limit
	*/
	KernelMemory *string

	/*
		Set meta data on a container
	*/
	Label []string

	/*
		Read in a line delimited file of labels
	*/
	LabelFile []string

	/*
		Add link to another container
	*/
	Link []string

	/*
		Container IPv4/IPv6 link-local addresses
	*/
	LinkLocalIp []string

	/*
		Logging driver for the container
	*/
	LogDriver *string

	/*
		Log driver options
	*/
	LogOpt []string

	/*
		Container MAC address (e.g., 92:d0:c6:0a:29:33)
	*/
	MacAddress *string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Memory soft limit
	*/
	MemoryReservation *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Tune container memory swappiness (0 to 100)
	*/
	MemorySwappiness *int64

	/*
		Attach a filesystem mount to the container
	*/
	Mount *string

	/*
		Assign a name to the container
	*/
	Name *string

	/*
		Connect a container to a network
	*/
	Net *string

	/*
		Add network-scoped alias for the container
	*/
	NetAlias []string

	/*
		Connect a container to a network
	*/
	Network *string

	/*
		Add network-scoped alias for the container
	*/
	NetworkAlias []string

	/*
		Disable any container-specified HEALTHCHECK
	*/
	NoHealthcheck *bool

	/*
		Disable OOM Killer
	*/
	OomKillDisable *bool

	/*
		Tune host's OOM preferences (-1000 to 1000)
	*/
	OomScoreAdj *int

	/*
		PID namespace to use
	*/
	Pid *string

	/*
		Tune container pids limit (set -1 for unlimited)
	*/
	PidsLimit *int64

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Give extended privileges to this container
	*/
	Privileged *bool

	/*
		Publish a container's port(s) to the host
	*/
	Publish []string

	/*
		Publish all exposed ports to random ports
	*/
	PublishAll *bool

	/*
		Pull image before creating ("always"|"missing"|"never")
	*/
	Pull *string

	/*
		Mount the container's root filesystem as read only
	*/
	ReadOnly *bool

	/*
		Restart policy to apply when a container exits
	*/
	Restart *string

	/*
		Automatically remove the container when it exits
	*/
	Rm *bool

	/*
		Runtime to use for this container
	*/
	Runtime *string

	/*
		Security Options
	*/
	SecurityOpt []string

	/*
		Size of /dev/shm
	*/
	ShmSize *string

	/*
		Signal to stop a container
	*/
	StopSignal *string

	/*
		Timeout (in seconds) to stop a container
	*/
	StopTimeout *int

	/*
		Storage driver options for the container
	*/
	StorageOpt []string

	/*
		Sysctl options
	*/
	Sysctl map[string]string

	/*
		Mount a tmpfs directory
	*/
	Tmpfs []string

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Ulimit options
	*/
	Ulimit *string

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		User namespace to use
	*/
	Userns *string

	/*
		UTS namespace to use
	*/
	Uts *string

	/*
		Bind mount a volume
	*/
	Volume []string

	/*
		Optional volume driver for the container
	*/
	VolumeDriver *string

	/*
		Mount volumes from the specified container(s)
	*/
	VolumesFrom []string

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerCreateCmd is wrapper of 'docker create'
------------------------------
create [OPTIONS] IMAGE [COMMAND] [ARG...]
Create a new container
------------------------------
*/
func DockerCreateCmd(opt DockerCreateOption, args []string) *exec.Cmd {
	cargs := []string{"create"}
	if opt.AddHost != nil {
		for _, str := range opt.AddHost {
			cargs = append(cargs, "--add-host")
			cargs = append(cargs, str)
		}
	}
	if opt.Attach != nil {
		for _, str := range opt.Attach {
			cargs = append(cargs, "--attach")
			cargs = append(cargs, str)
		}
	}
	if opt.BlkioWeight != nil {
		cargs = append(cargs, "--blkio-weight="+fmt.Sprint(*opt.BlkioWeight))
	}
	if opt.BlkioWeightDevice != nil {
		for _, str := range opt.BlkioWeightDevice {
			cargs = append(cargs, "--blkio-weight-device")
			cargs = append(cargs, str)
		}
	}
	if opt.CapAdd != nil {
		for _, str := range opt.CapAdd {
			cargs = append(cargs, "--cap-add")
			cargs = append(cargs, str)
		}
	}
	if opt.CapDrop != nil {
		for _, str := range opt.CapDrop {
			cargs = append(cargs, "--cap-drop")
			cargs = append(cargs, str)
		}
	}
	if opt.CgroupParent != nil {
		cargs = append(cargs, "--cgroup-parent="+fmt.Sprint(*opt.CgroupParent))
	}
	if opt.Cgroupns != nil {
		cargs = append(cargs, "--cgroupns="+fmt.Sprint(*opt.Cgroupns))
	}
	if opt.Cidfile != nil {
		cargs = append(cargs, "--cidfile="+fmt.Sprint(*opt.Cidfile))
	}
	if opt.CpuCount != nil {
		cargs = append(cargs, "--cpu-count="+fmt.Sprint(*opt.CpuCount))
	}
	if opt.CpuPercent != nil {
		cargs = append(cargs, "--cpu-percent="+fmt.Sprint(*opt.CpuPercent))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuRtPeriod != nil {
		cargs = append(cargs, "--cpu-rt-period="+fmt.Sprint(*opt.CpuRtPeriod))
	}
	if opt.CpuRtRuntime != nil {
		cargs = append(cargs, "--cpu-rt-runtime="+fmt.Sprint(*opt.CpuRtRuntime))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.Cpus != nil {
		cargs = append(cargs, "--cpus="+fmt.Sprint(*opt.Cpus))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.Device != nil {
		for _, str := range opt.Device {
			cargs = append(cargs, "--device")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceCgroupRule != nil {
		for _, str := range opt.DeviceCgroupRule {
			cargs = append(cargs, "--device-cgroup-rule")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadBps != nil {
		for _, str := range opt.DeviceReadBps {
			cargs = append(cargs, "--device-read-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadIops != nil {
		for _, str := range opt.DeviceReadIops {
			cargs = append(cargs, "--device-read-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteBps != nil {
		for _, str := range opt.DeviceWriteBps {
			cargs = append(cargs, "--device-write-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteIops != nil {
		for _, str := range opt.DeviceWriteIops {
			cargs = append(cargs, "--device-write-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Dns != nil {
		for _, str := range opt.Dns {
			cargs = append(cargs, "--dns")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOpt != nil {
		for _, str := range opt.DnsOpt {
			cargs = append(cargs, "--dns-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOption != nil {
		for _, str := range opt.DnsOption {
			cargs = append(cargs, "--dns-option")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsSearch != nil {
		for _, str := range opt.DnsSearch {
			cargs = append(cargs, "--dns-search")
			cargs = append(cargs, str)
		}
	}
	if opt.Domainname != nil {
		cargs = append(cargs, "--domainname="+fmt.Sprint(*opt.Domainname))
	}
	if opt.Entrypoint != nil {
		cargs = append(cargs, "--entrypoint="+fmt.Sprint(*opt.Entrypoint))
	}
	if opt.Env != nil {
		for _, str := range opt.Env {
			cargs = append(cargs, "--env")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvFile != nil {
		for _, str := range opt.EnvFile {
			cargs = append(cargs, "--env-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Expose != nil {
		for _, str := range opt.Expose {
			cargs = append(cargs, "--expose")
			cargs = append(cargs, str)
		}
	}
	if opt.Gpus != nil {
		cargs = append(cargs, "--gpus="+fmt.Sprint(*opt.Gpus))
	}
	if opt.GroupAdd != nil {
		for _, str := range opt.GroupAdd {
			cargs = append(cargs, "--group-add")
			cargs = append(cargs, str)
		}
	}
	if opt.HealthCmd != nil {
		cargs = append(cargs, "--health-cmd="+fmt.Sprint(*opt.HealthCmd))
	}
	if opt.HealthInterval != nil {
		cargs = append(cargs, "--health-interval="+fmt.Sprint(*opt.HealthInterval))
	}
	if opt.HealthRetries != nil {
		cargs = append(cargs, "--health-retries="+fmt.Sprint(*opt.HealthRetries))
	}
	if opt.HealthStartPeriod != nil {
		cargs = append(cargs, "--health-start-period="+fmt.Sprint(*opt.HealthStartPeriod))
	}
	if opt.HealthTimeout != nil {
		cargs = append(cargs, "--health-timeout="+fmt.Sprint(*opt.HealthTimeout))
	}
	if opt.Help != nil {
		cargs = append(cargs, "--help="+fmt.Sprint(*opt.Help))
	}
	if opt.Hostname != nil {
		cargs = append(cargs, "--hostname="+fmt.Sprint(*opt.Hostname))
	}
	if opt.Init != nil {
		cargs = append(cargs, "--init="+fmt.Sprint(*opt.Init))
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	if opt.IoMaxbandwidth != nil {
		cargs = append(cargs, "--io-maxbandwidth="+fmt.Sprint(*opt.IoMaxbandwidth))
	}
	if opt.IoMaxiops != nil {
		cargs = append(cargs, "--io-maxiops="+fmt.Sprint(*opt.IoMaxiops))
	}
	if opt.Ip != nil {
		cargs = append(cargs, "--ip="+fmt.Sprint(*opt.Ip))
	}
	if opt.Ip6 != nil {
		cargs = append(cargs, "--ip6="+fmt.Sprint(*opt.Ip6))
	}
	if opt.Ipc != nil {
		cargs = append(cargs, "--ipc="+fmt.Sprint(*opt.Ipc))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.KernelMemory != nil {
		cargs = append(cargs, "--kernel-memory="+fmt.Sprint(*opt.KernelMemory))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.LabelFile != nil {
		for _, str := range opt.LabelFile {
			cargs = append(cargs, "--label-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Link != nil {
		for _, str := range opt.Link {
			cargs = append(cargs, "--link")
			cargs = append(cargs, str)
		}
	}
	if opt.LinkLocalIp != nil {
		for _, str := range opt.LinkLocalIp {
			cargs = append(cargs, "--link-local-ip")
			cargs = append(cargs, str)
		}
	}
	if opt.LogDriver != nil {
		cargs = append(cargs, "--log-driver="+fmt.Sprint(*opt.LogDriver))
	}
	if opt.LogOpt != nil {
		for _, str := range opt.LogOpt {
			cargs = append(cargs, "--log-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.MacAddress != nil {
		cargs = append(cargs, "--mac-address="+fmt.Sprint(*opt.MacAddress))
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemoryReservation != nil {
		cargs = append(cargs, "--memory-reservation="+fmt.Sprint(*opt.MemoryReservation))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.MemorySwappiness != nil {
		cargs = append(cargs, "--memory-swappiness="+fmt.Sprint(*opt.MemorySwappiness))
	}
	if opt.Mount != nil {
		cargs = append(cargs, "--mount="+fmt.Sprint(*opt.Mount))
	}
	if opt.Name != nil {
		cargs = append(cargs, "--name="+fmt.Sprint(*opt.Name))
	}
	if opt.Net != nil {
		cargs = append(cargs, "--net="+fmt.Sprint(*opt.Net))
	}
	if opt.NetAlias != nil {
		for _, str := range opt.NetAlias {
			cargs = append(cargs, "--net-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NetworkAlias != nil {
		for _, str := range opt.NetworkAlias {
			cargs = append(cargs, "--network-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.NoHealthcheck != nil {
		cargs = append(cargs, "--no-healthcheck="+fmt.Sprint(*opt.NoHealthcheck))
	}
	if opt.OomKillDisable != nil {
		cargs = append(cargs, "--oom-kill-disable="+fmt.Sprint(*opt.OomKillDisable))
	}
	if opt.OomScoreAdj != nil {
		cargs = append(cargs, "--oom-score-adj="+fmt.Sprint(*opt.OomScoreAdj))
	}
	if opt.Pid != nil {
		cargs = append(cargs, "--pid="+fmt.Sprint(*opt.Pid))
	}
	if opt.PidsLimit != nil {
		cargs = append(cargs, "--pids-limit="+fmt.Sprint(*opt.PidsLimit))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Privileged != nil {
		cargs = append(cargs, "--privileged="+fmt.Sprint(*opt.Privileged))
	}
	if opt.Publish != nil {
		for _, str := range opt.Publish {
			cargs = append(cargs, "--publish")
			cargs = append(cargs, str)
		}
	}
	if opt.PublishAll != nil {
		cargs = append(cargs, "--publish-all="+fmt.Sprint(*opt.PublishAll))
	}
	if opt.Pull != nil {
		cargs = append(cargs, "--pull="+fmt.Sprint(*opt.Pull))
	}
	if opt.ReadOnly != nil {
		cargs = append(cargs, "--read-only="+fmt.Sprint(*opt.ReadOnly))
	}
	if opt.Restart != nil {
		cargs = append(cargs, "--restart="+fmt.Sprint(*opt.Restart))
	}
	if opt.Rm != nil {
		cargs = append(cargs, "--rm="+fmt.Sprint(*opt.Rm))
	}
	if opt.Runtime != nil {
		cargs = append(cargs, "--runtime="+fmt.Sprint(*opt.Runtime))
	}
	if opt.SecurityOpt != nil {
		for _, str := range opt.SecurityOpt {
			cargs = append(cargs, "--security-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.ShmSize != nil {
		cargs = append(cargs, "--shm-size="+fmt.Sprint(*opt.ShmSize))
	}
	if opt.StopSignal != nil {
		cargs = append(cargs, "--stop-signal="+fmt.Sprint(*opt.StopSignal))
	}
	if opt.StopTimeout != nil {
		cargs = append(cargs, "--stop-timeout="+fmt.Sprint(*opt.StopTimeout))
	}
	if opt.StorageOpt != nil {
		for _, str := range opt.StorageOpt {
			cargs = append(cargs, "--storage-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.Sysctl != nil {
		for key, val := range opt.Sysctl {
			cargs = append(cargs, "--sysctl")
			cargs = append(cargs, key+"="+val)
		}
	}
	if opt.Tmpfs != nil {
		for _, str := range opt.Tmpfs {
			cargs = append(cargs, "--tmpfs")
			cargs = append(cargs, str)
		}
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.Userns != nil {
		cargs = append(cargs, "--userns="+fmt.Sprint(*opt.Userns))
	}
	if opt.Uts != nil {
		cargs = append(cargs, "--uts="+fmt.Sprint(*opt.Uts))
	}
	if opt.Volume != nil {
		for _, str := range opt.Volume {
			cargs = append(cargs, "--volume")
			cargs = append(cargs, str)
		}
	}
	if opt.VolumeDriver != nil {
		cargs = append(cargs, "--volume-driver="+fmt.Sprint(*opt.VolumeDriver))
	}
	if opt.VolumesFrom != nil {
		for _, str := range opt.VolumesFrom {
			cargs = append(cargs, "--volumes-from")
			cargs = append(cargs, str)
		}
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerDiffCmd is wrapper of 'docker diff'
------------------------------
diff CONTAINER
Inspect changes to files or directories on a container's filesystem
------------------------------
*/
func DockerDiffCmd(args []string) *exec.Cmd {
	cargs := []string{"diff"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerEventsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Show all events created since timestamp
	*/
	Since *string

	/*
		Stream events until this timestamp
	*/
	Until *string
}

/*
DockerEventsCmd is wrapper of 'docker events'
------------------------------
events [OPTIONS]
Get real time events from the server
------------------------------
*/
func DockerEventsCmd(opt DockerEventsOption, args []string) *exec.Cmd {
	cargs := []string{"events"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Since != nil {
		cargs = append(cargs, "--since="+fmt.Sprint(*opt.Since))
	}
	if opt.Until != nil {
		cargs = append(cargs, "--until="+fmt.Sprint(*opt.Until))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerExecOption struct {
	/*
		Detached mode: run command in the background
	*/
	Detach *bool

	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Set environment variables
	*/
	Env []string

	/*
		Read in a file of environment variables
	*/
	EnvFile []string

	/*
		Keep STDIN open even if not attached
	*/
	Interactive *bool

	/*
		Give extended privileges to the command
	*/
	Privileged *bool

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerExecCmd is wrapper of 'docker exec'
------------------------------
exec [OPTIONS] CONTAINER COMMAND [ARG...]
Run a command in a running container
------------------------------
*/
func DockerExecCmd(opt DockerExecOption, args []string) *exec.Cmd {
	cargs := []string{"exec"}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.Env != nil {
		for _, str := range opt.Env {
			cargs = append(cargs, "--env")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvFile != nil {
		for _, str := range opt.EnvFile {
			cargs = append(cargs, "--env-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	if opt.Privileged != nil {
		cargs = append(cargs, "--privileged="+fmt.Sprint(*opt.Privileged))
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerExportOption struct {
	/*
		Write to a file, instead of STDOUT
	*/
	Output *string
}

/*
DockerExportCmd is wrapper of 'docker export'
------------------------------
export [OPTIONS] CONTAINER
Export a container's filesystem as a tar archive
------------------------------
*/
func DockerExportCmd(opt DockerExportOption, args []string) *exec.Cmd {
	cargs := []string{"export"}
	if opt.Output != nil {
		cargs = append(cargs, "--output="+fmt.Sprint(*opt.Output))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerHistoryOption struct {
	/*
		Pretty-print images using a Go template
	*/
	Format *string

	/*
		Print sizes and dates in human readable format
	*/
	Human *bool

	/*
		Don't truncate output
	*/
	NoTrunc *bool

	/*
		Only show image IDs
	*/
	Quiet *bool
}

/*
DockerHistoryCmd is wrapper of 'docker history'
------------------------------
history [OPTIONS] IMAGE
Show the history of an image
------------------------------
*/
func DockerHistoryCmd(opt DockerHistoryOption, args []string) *exec.Cmd {
	cargs := []string{"history"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Human != nil {
		cargs = append(cargs, "--human="+fmt.Sprint(*opt.Human))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerImageCmd is wrapper of 'docker image'
------------------------------
image
Manage images
------------------------------
*/
func DockerImageCmd(args []string) *exec.Cmd {
	cargs := []string{"image"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageBuildOption struct {
	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	AddHost []string

	/*
		Set build-time variables
	*/
	BuildArg []string

	/*
		Images to consider as cache sources
	*/
	CacheFrom *string

	/*
		Optional parent cgroup for the container
	*/
	CgroupParent *string

	/*
		Compress the build context using gzip
	*/
	Compress *bool

	/*
		Limit the CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit the CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Name of the Dockerfile (Default is 'PATH/Dockerfile')
	*/
	File *string

	/*
		Always remove intermediate containers
	*/
	ForceRm *bool

	/*
		Write the image ID to the file
	*/
	Iidfile *string

	/*
		Container isolation technology
	*/
	Isolation *string

	/*
		Set metadata for an image
	*/
	Label []string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Set the networking mode for the RUN instructions during build
	*/
	Network *string

	/*
		Do not use cache when building the image
	*/
	NoCache *bool

	/*
		Output destination (format: type=local,dest=path)
	*/
	Output *string

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Set type of progress output (auto, plain, tty). Use plain to show container output
	*/
	Progress *string

	/*
		Always attempt to pull a newer version of the image
	*/
	Pull *bool

	/*
		Suppress the build output and print image ID on success
	*/
	Quiet *bool

	/*
		Remove intermediate containers after a successful build
	*/
	Rm *bool

	/*
		Secret file to expose to the build (only if BuildKit enabled): id=mysecret,src=/local/secret
	*/
	Secret *string

	/*
		Security options
	*/
	SecurityOpt *string

	/*
		Size of /dev/shm
	*/
	ShmSize *string

	/*
		Squash newly built layers into a single new layer
	*/
	Squash *bool

	/*
		SSH agent socket or keys to expose to the build (only if BuildKit enabled) (format: default|<id>[=<socket>|<key>[,<key>]])
	*/
	Ssh *string

	/*
		Stream attaches to server to negotiate build context
	*/
	Stream *bool

	/*
		Name and optionally a tag in the 'name:tag' format
	*/
	Tag []string

	/*
		Set the target build stage to build.
	*/
	Target *string

	/*
		Ulimit options
	*/
	Ulimit *string
}

/*
DockerImageBuildCmd is wrapper of 'docker image build'
------------------------------
build [OPTIONS] PATH | URL | -
Build an image from a Dockerfile
------------------------------
*/
func DockerImageBuildCmd(opt DockerImageBuildOption, args []string) *exec.Cmd {
	cargs := []string{"image", "build"}
	if opt.AddHost != nil {
		for _, str := range opt.AddHost {
			cargs = append(cargs, "--add-host")
			cargs = append(cargs, str)
		}
	}
	if opt.BuildArg != nil {
		for _, str := range opt.BuildArg {
			cargs = append(cargs, "--build-arg")
			cargs = append(cargs, str)
		}
	}
	if opt.CacheFrom != nil {
		cargs = append(cargs, "--cache-from="+fmt.Sprint(*opt.CacheFrom))
	}
	if opt.CgroupParent != nil {
		cargs = append(cargs, "--cgroup-parent="+fmt.Sprint(*opt.CgroupParent))
	}
	if opt.Compress != nil {
		cargs = append(cargs, "--compress="+fmt.Sprint(*opt.Compress))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.File != nil {
		cargs = append(cargs, "--file="+fmt.Sprint(*opt.File))
	}
	if opt.ForceRm != nil {
		cargs = append(cargs, "--force-rm="+fmt.Sprint(*opt.ForceRm))
	}
	if opt.Iidfile != nil {
		cargs = append(cargs, "--iidfile="+fmt.Sprint(*opt.Iidfile))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NoCache != nil {
		cargs = append(cargs, "--no-cache="+fmt.Sprint(*opt.NoCache))
	}
	if opt.Output != nil {
		cargs = append(cargs, "--output="+fmt.Sprint(*opt.Output))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Progress != nil {
		cargs = append(cargs, "--progress="+fmt.Sprint(*opt.Progress))
	}
	if opt.Pull != nil {
		cargs = append(cargs, "--pull="+fmt.Sprint(*opt.Pull))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Rm != nil {
		cargs = append(cargs, "--rm="+fmt.Sprint(*opt.Rm))
	}
	if opt.Secret != nil {
		cargs = append(cargs, "--secret="+fmt.Sprint(*opt.Secret))
	}
	if opt.SecurityOpt != nil {
		cargs = append(cargs, "--security-opt="+fmt.Sprint(*opt.SecurityOpt))
	}
	if opt.ShmSize != nil {
		cargs = append(cargs, "--shm-size="+fmt.Sprint(*opt.ShmSize))
	}
	if opt.Squash != nil {
		cargs = append(cargs, "--squash="+fmt.Sprint(*opt.Squash))
	}
	if opt.Ssh != nil {
		cargs = append(cargs, "--ssh="+fmt.Sprint(*opt.Ssh))
	}
	if opt.Stream != nil {
		cargs = append(cargs, "--stream="+fmt.Sprint(*opt.Stream))
	}
	if opt.Tag != nil {
		for _, str := range opt.Tag {
			cargs = append(cargs, "--tag")
			cargs = append(cargs, str)
		}
	}
	if opt.Target != nil {
		cargs = append(cargs, "--target="+fmt.Sprint(*opt.Target))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageHistoryOption struct {
	/*
		Pretty-print images using a Go template
	*/
	Format *string

	/*
		Print sizes and dates in human readable format
	*/
	Human *bool

	/*
		Don't truncate output
	*/
	NoTrunc *bool

	/*
		Only show image IDs
	*/
	Quiet *bool
}

/*
DockerImageHistoryCmd is wrapper of 'docker image history'
------------------------------
history [OPTIONS] IMAGE
Show the history of an image
------------------------------
*/
func DockerImageHistoryCmd(opt DockerImageHistoryOption, args []string) *exec.Cmd {
	cargs := []string{"image", "history"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Human != nil {
		cargs = append(cargs, "--human="+fmt.Sprint(*opt.Human))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageImportOption struct {
	/*
		Apply Dockerfile instruction to the created image
	*/
	Change []string

	/*
		Set commit message for imported image
	*/
	Message *string

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string
}

/*
DockerImageImportCmd is wrapper of 'docker image import'
------------------------------
import [OPTIONS] file|URL|- [REPOSITORY[:TAG]]
Import the contents from a tarball to create a filesystem image
------------------------------
*/
func DockerImageImportCmd(opt DockerImageImportOption, args []string) *exec.Cmd {
	cargs := []string{"image", "import"}
	if opt.Change != nil {
		for _, str := range opt.Change {
			cargs = append(cargs, "--change")
			cargs = append(cargs, str)
		}
	}
	if opt.Message != nil {
		cargs = append(cargs, "--message="+fmt.Sprint(*opt.Message))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string
}

/*
DockerImageInspectCmd is wrapper of 'docker image inspect'
------------------------------
inspect [OPTIONS] IMAGE [IMAGE...]
Display detailed information on one or more images
------------------------------
*/
func DockerImageInspectCmd(opt DockerImageInspectOption, args []string) *exec.Cmd {
	cargs := []string{"image", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageLoadOption struct {
	/*
		Read from tar archive file, instead of STDIN
	*/
	Input *string

	/*
		Suppress the load output
	*/
	Quiet *bool
}

/*
DockerImageLoadCmd is wrapper of 'docker image load'
------------------------------
load [OPTIONS]
Load an image from a tar archive or STDIN
------------------------------
*/
func DockerImageLoadCmd(opt DockerImageLoadOption, args []string) *exec.Cmd {
	cargs := []string{"image", "load"}
	if opt.Input != nil {
		cargs = append(cargs, "--input="+fmt.Sprint(*opt.Input))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageLsOption struct {
	/*
		Show all images (default hides intermediate images)
	*/
	All *bool

	/*
		Show digests
	*/
	Digests *bool

	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print images using a Go template
	*/
	Format *string

	/*
		Don't truncate output
	*/
	NoTrunc *bool

	/*
		Only show image IDs
	*/
	Quiet *bool
}

/*
DockerImageLsCmd is wrapper of 'docker image ls'
------------------------------
ls [OPTIONS] [REPOSITORY[:TAG]]
List images
------------------------------
*/
func DockerImageLsCmd(opt DockerImageLsOption, args []string) *exec.Cmd {
	cargs := []string{"image", "ls"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Digests != nil {
		cargs = append(cargs, "--digests="+fmt.Sprint(*opt.Digests))
	}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImagePruneOption struct {
	/*
		Remove all unused images, not just dangling ones
	*/
	All *bool

	/*
		Provide filter values (e.g. 'until=<timestamp>')
	*/
	Filter *string

	/*
		Do not prompt for confirmation
	*/
	Force *bool
}

/*
DockerImagePruneCmd is wrapper of 'docker image prune'
------------------------------
prune [OPTIONS]
Remove unused images
------------------------------
*/
func DockerImagePruneCmd(opt DockerImagePruneOption, args []string) *exec.Cmd {
	cargs := []string{"image", "prune"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImagePullOption struct {
	/*
		Download all tagged images in the repository
	*/
	AllTags *bool

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Suppress verbose output
	*/
	Quiet *bool
}

/*
DockerImagePullCmd is wrapper of 'docker image pull'
------------------------------
pull [OPTIONS] NAME[:TAG|@DIGEST]
Pull an image or a repository from a registry
------------------------------
*/
func DockerImagePullCmd(opt DockerImagePullOption, args []string) *exec.Cmd {
	cargs := []string{"image", "pull"}
	if opt.AllTags != nil {
		cargs = append(cargs, "--all-tags="+fmt.Sprint(*opt.AllTags))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImagePushOption struct {
	/*
		Push all tagged images in the repository
	*/
	AllTags *bool

	/*
		Skip image signing
	*/
	DisableContentTrust *bool

	/*
		Suppress verbose output
	*/
	Quiet *bool
}

/*
DockerImagePushCmd is wrapper of 'docker image push'
------------------------------
push [OPTIONS] NAME[:TAG]
Push an image or a repository to a registry
------------------------------
*/
func DockerImagePushCmd(opt DockerImagePushOption, args []string) *exec.Cmd {
	cargs := []string{"image", "push"}
	if opt.AllTags != nil {
		cargs = append(cargs, "--all-tags="+fmt.Sprint(*opt.AllTags))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageRmOption struct {
	/*
		Force removal of the image
	*/
	Force *bool

	/*
		Do not delete untagged parents
	*/
	NoPrune *bool
}

/*
DockerImageRmCmd is wrapper of 'docker image rm'
------------------------------
rm [OPTIONS] IMAGE [IMAGE...]
Remove one or more images
------------------------------
*/
func DockerImageRmCmd(opt DockerImageRmOption, args []string) *exec.Cmd {
	cargs := []string{"image", "rm"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	if opt.NoPrune != nil {
		cargs = append(cargs, "--no-prune="+fmt.Sprint(*opt.NoPrune))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImageSaveOption struct {
	/*
		Write to a file, instead of STDOUT
	*/
	Output *string
}

/*
DockerImageSaveCmd is wrapper of 'docker image save'
------------------------------
save [OPTIONS] IMAGE [IMAGE...]
Save one or more images to a tar archive (streamed to STDOUT by default)
------------------------------
*/
func DockerImageSaveCmd(opt DockerImageSaveOption, args []string) *exec.Cmd {
	cargs := []string{"image", "save"}
	if opt.Output != nil {
		cargs = append(cargs, "--output="+fmt.Sprint(*opt.Output))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerImageTagCmd is wrapper of 'docker image tag'
------------------------------
tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
------------------------------
*/
func DockerImageTagCmd(args []string) *exec.Cmd {
	cargs := []string{"image", "tag"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImagesOption struct {
	/*
		Show all images (default hides intermediate images)
	*/
	All *bool

	/*
		Show digests
	*/
	Digests *bool

	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print images using a Go template
	*/
	Format *string

	/*
		Don't truncate output
	*/
	NoTrunc *bool

	/*
		Only show image IDs
	*/
	Quiet *bool
}

/*
DockerImagesCmd is wrapper of 'docker images'
------------------------------
images [OPTIONS] [REPOSITORY[:TAG]]
List images
------------------------------
*/
func DockerImagesCmd(opt DockerImagesOption, args []string) *exec.Cmd {
	cargs := []string{"images"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Digests != nil {
		cargs = append(cargs, "--digests="+fmt.Sprint(*opt.Digests))
	}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerImportOption struct {
	/*
		Apply Dockerfile instruction to the created image
	*/
	Change []string

	/*
		Set commit message for imported image
	*/
	Message *string

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string
}

/*
DockerImportCmd is wrapper of 'docker import'
------------------------------
import [OPTIONS] file|URL|- [REPOSITORY[:TAG]]
Import the contents from a tarball to create a filesystem image
------------------------------
*/
func DockerImportCmd(opt DockerImportOption, args []string) *exec.Cmd {
	cargs := []string{"import"}
	if opt.Change != nil {
		for _, str := range opt.Change {
			cargs = append(cargs, "--change")
			cargs = append(cargs, str)
		}
	}
	if opt.Message != nil {
		cargs = append(cargs, "--message="+fmt.Sprint(*opt.Message))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerInfoOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string
}

/*
DockerInfoCmd is wrapper of 'docker info'
------------------------------
info [OPTIONS]
Display system-wide information
------------------------------
*/
func DockerInfoCmd(opt DockerInfoOption, args []string) *exec.Cmd {
	cargs := []string{"info"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Display total file sizes if the type is container
	*/
	Size *bool

	/*
		Return JSON for specified type
	*/
	Type *string
}

/*
DockerInspectCmd is wrapper of 'docker inspect'
------------------------------
inspect [OPTIONS] NAME|ID [NAME|ID...]
Return low-level information on Docker objects
------------------------------
*/
func DockerInspectCmd(opt DockerInspectOption, args []string) *exec.Cmd {
	cargs := []string{"inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Size != nil {
		cargs = append(cargs, "--size="+fmt.Sprint(*opt.Size))
	}
	if opt.Type != nil {
		cargs = append(cargs, "--type="+fmt.Sprint(*opt.Type))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerKillOption struct {
	/*
		Signal to send to the container
	*/
	Signal *string
}

/*
DockerKillCmd is wrapper of 'docker kill'
------------------------------
kill [OPTIONS] CONTAINER [CONTAINER...]
Kill one or more running containers
------------------------------
*/
func DockerKillCmd(opt DockerKillOption, args []string) *exec.Cmd {
	cargs := []string{"kill"}
	if opt.Signal != nil {
		cargs = append(cargs, "--signal="+fmt.Sprint(*opt.Signal))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerLoadOption struct {
	/*
		Read from tar archive file, instead of STDIN
	*/
	Input *string

	/*
		Suppress the load output
	*/
	Quiet *bool
}

/*
DockerLoadCmd is wrapper of 'docker load'
------------------------------
load [OPTIONS]
Load an image from a tar archive or STDIN
------------------------------
*/
func DockerLoadCmd(opt DockerLoadOption, args []string) *exec.Cmd {
	cargs := []string{"load"}
	if opt.Input != nil {
		cargs = append(cargs, "--input="+fmt.Sprint(*opt.Input))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerLoginOption struct {
	/*
		Password
	*/
	Password *string

	/*
		Take the password from stdin
	*/
	PasswordStdin *bool

	/*
		Username
	*/
	Username *string
}

/*
DockerLoginCmd is wrapper of 'docker login'
------------------------------
login [OPTIONS] [SERVER]
Log in to a Docker registry
------------------------------
*/
func DockerLoginCmd(opt DockerLoginOption, args []string) *exec.Cmd {
	cargs := []string{"login"}
	if opt.Password != nil {
		cargs = append(cargs, "--password="+fmt.Sprint(*opt.Password))
	}
	if opt.PasswordStdin != nil {
		cargs = append(cargs, "--password-stdin="+fmt.Sprint(*opt.PasswordStdin))
	}
	if opt.Username != nil {
		cargs = append(cargs, "--username="+fmt.Sprint(*opt.Username))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerLogoutCmd is wrapper of 'docker logout'
------------------------------
logout [SERVER]
Log out from a Docker registry
------------------------------
*/
func DockerLogoutCmd(args []string) *exec.Cmd {
	cargs := []string{"logout"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerLogsOption struct {
	/*
		Show extra details provided to logs
	*/
	Details *bool

	/*
		Follow log output
	*/
	Follow *bool

	/*
		Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
	*/
	Since *string

	/*
		Number of lines to show from the end of the logs
	*/
	Tail *string

	/*
		Show timestamps
	*/
	Timestamps *bool

	/*
		Show logs before a timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
	*/
	Until *string
}

/*
DockerLogsCmd is wrapper of 'docker logs'
------------------------------
logs [OPTIONS] CONTAINER
Fetch the logs of a container
------------------------------
*/
func DockerLogsCmd(opt DockerLogsOption, args []string) *exec.Cmd {
	cargs := []string{"logs"}
	if opt.Details != nil {
		cargs = append(cargs, "--details="+fmt.Sprint(*opt.Details))
	}
	if opt.Follow != nil {
		cargs = append(cargs, "--follow="+fmt.Sprint(*opt.Follow))
	}
	if opt.Since != nil {
		cargs = append(cargs, "--since="+fmt.Sprint(*opt.Since))
	}
	if opt.Tail != nil {
		cargs = append(cargs, "--tail="+fmt.Sprint(*opt.Tail))
	}
	if opt.Timestamps != nil {
		cargs = append(cargs, "--timestamps="+fmt.Sprint(*opt.Timestamps))
	}
	if opt.Until != nil {
		cargs = append(cargs, "--until="+fmt.Sprint(*opt.Until))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerManifestCmd is wrapper of 'docker manifest'
------------------------------
manifest COMMAND
Manage Docker image manifests and manifest lists
------------------------------
*/
func DockerManifestCmd(args []string) *exec.Cmd {
	cargs := []string{"manifest"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerManifestAnnotateOption struct {
	/*
		Set architecture
	*/
	Arch *string

	/*
		Set operating system
	*/
	Os *string

	/*
		Set operating system feature
	*/
	OsFeatures *string

	/*
		Set operating system version
	*/
	OsVersion *string

	/*
		Set architecture variant
	*/
	Variant *string
}

/*
DockerManifestAnnotateCmd is wrapper of 'docker manifest annotate'
------------------------------
annotate [OPTIONS] MANIFEST_LIST MANIFEST
Add additional information to a local image manifest
------------------------------
*/
func DockerManifestAnnotateCmd(opt DockerManifestAnnotateOption, args []string) *exec.Cmd {
	cargs := []string{"manifest", "annotate"}
	if opt.Arch != nil {
		cargs = append(cargs, "--arch="+fmt.Sprint(*opt.Arch))
	}
	if opt.Os != nil {
		cargs = append(cargs, "--os="+fmt.Sprint(*opt.Os))
	}
	if opt.OsFeatures != nil {
		cargs = append(cargs, "--os-features="+fmt.Sprint(*opt.OsFeatures))
	}
	if opt.OsVersion != nil {
		cargs = append(cargs, "--os-version="+fmt.Sprint(*opt.OsVersion))
	}
	if opt.Variant != nil {
		cargs = append(cargs, "--variant="+fmt.Sprint(*opt.Variant))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerManifestCreateOption struct {
	/*
		Amend an existing manifest list
	*/
	Amend *bool

	/*
		Allow communication with an insecure registry
	*/
	Insecure *bool
}

/*
DockerManifestCreateCmd is wrapper of 'docker manifest create'
------------------------------
create MANIFEST_LIST MANIFEST [MANIFEST...]
Create a local manifest list for annotating and pushing to a registry
------------------------------
*/
func DockerManifestCreateCmd(opt DockerManifestCreateOption, args []string) *exec.Cmd {
	cargs := []string{"manifest", "create"}
	if opt.Amend != nil {
		cargs = append(cargs, "--amend="+fmt.Sprint(*opt.Amend))
	}
	if opt.Insecure != nil {
		cargs = append(cargs, "--insecure="+fmt.Sprint(*opt.Insecure))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerManifestInspectOption struct {
	/*
		Allow communication with an insecure registry
	*/
	Insecure *bool

	/*
		Output additional info including layers and platform
	*/
	Verbose *bool
}

/*
DockerManifestInspectCmd is wrapper of 'docker manifest inspect'
------------------------------
inspect [OPTIONS] [MANIFEST_LIST] MANIFEST
Display an image manifest, or manifest list
------------------------------
*/
func DockerManifestInspectCmd(opt DockerManifestInspectOption, args []string) *exec.Cmd {
	cargs := []string{"manifest", "inspect"}
	if opt.Insecure != nil {
		cargs = append(cargs, "--insecure="+fmt.Sprint(*opt.Insecure))
	}
	if opt.Verbose != nil {
		cargs = append(cargs, "--verbose="+fmt.Sprint(*opt.Verbose))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerManifestPushOption struct {
	/*
		Allow push to an insecure registry
	*/
	Insecure *bool

	/*
		Remove the local manifest list after push
	*/
	Purge *bool
}

/*
DockerManifestPushCmd is wrapper of 'docker manifest push'
------------------------------
push [OPTIONS] MANIFEST_LIST
Push a manifest list to a repository
------------------------------
*/
func DockerManifestPushCmd(opt DockerManifestPushOption, args []string) *exec.Cmd {
	cargs := []string{"manifest", "push"}
	if opt.Insecure != nil {
		cargs = append(cargs, "--insecure="+fmt.Sprint(*opt.Insecure))
	}
	if opt.Purge != nil {
		cargs = append(cargs, "--purge="+fmt.Sprint(*opt.Purge))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerManifestRmCmd is wrapper of 'docker manifest rm'
------------------------------
rm MANIFEST_LIST [MANIFEST_LIST...]
Delete one or more manifest lists from local storage
------------------------------
*/
func DockerManifestRmCmd(args []string) *exec.Cmd {
	cargs := []string{"manifest", "rm"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerNetworkCmd is wrapper of 'docker network'
------------------------------
network
Manage networks
------------------------------
*/
func DockerNetworkCmd(args []string) *exec.Cmd {
	cargs := []string{"network"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNetworkConnectOption struct {
	/*
		Add network-scoped alias for the container
	*/
	Alias *string

	/*
		driver options for the network
	*/
	DriverOpt *string

	/*
		IPv4 address (e.g., 172.30.100.104)
	*/
	Ip *string

	/*
		IPv6 address (e.g., 2001:db8::33)
	*/
	Ip6 *string

	/*
		Add link to another container
	*/
	Link []string

	/*
		Add a link-local address for the container
	*/
	LinkLocalIp *string
}

/*
DockerNetworkConnectCmd is wrapper of 'docker network connect'
------------------------------
connect [OPTIONS] NETWORK CONTAINER
Connect a container to a network
------------------------------
*/
func DockerNetworkConnectCmd(opt DockerNetworkConnectOption, args []string) *exec.Cmd {
	cargs := []string{"network", "connect"}
	if opt.Alias != nil {
		cargs = append(cargs, "--alias="+fmt.Sprint(*opt.Alias))
	}
	if opt.DriverOpt != nil {
		cargs = append(cargs, "--driver-opt="+fmt.Sprint(*opt.DriverOpt))
	}
	if opt.Ip != nil {
		cargs = append(cargs, "--ip="+fmt.Sprint(*opt.Ip))
	}
	if opt.Ip6 != nil {
		cargs = append(cargs, "--ip6="+fmt.Sprint(*opt.Ip6))
	}
	if opt.Link != nil {
		for _, str := range opt.Link {
			cargs = append(cargs, "--link")
			cargs = append(cargs, str)
		}
	}
	if opt.LinkLocalIp != nil {
		cargs = append(cargs, "--link-local-ip="+fmt.Sprint(*opt.LinkLocalIp))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNetworkCreateOption struct {
	/*
		Enable manual container attachment
	*/
	Attachable *bool

	/*
		Auxiliary IPv4 or IPv6 addresses used by Network driver
	*/
	AuxAddress map[string]string

	/*
		The network from which to copy the configuration
	*/
	ConfigFrom *string

	/*
		Create a configuration only network
	*/
	ConfigOnly *bool

	/*
		Driver to manage the Network
	*/
	Driver *string

	/*
		IPv4 or IPv6 Gateway for the master subnet
	*/
	Gateway *string

	/*
		Create swarm routing-mesh network
	*/
	Ingress *bool

	/*
		Restrict external access to the network
	*/
	Internal *bool

	/*
		Allocate container ip from a sub-range
	*/
	IpRange *string

	/*
		IP Address Management Driver
	*/
	IpamDriver *string

	/*
		Set IPAM driver specific options
	*/
	IpamOpt map[string]string

	/*
		Enable IPv6 networking
	*/
	Ipv6 *bool

	/*
		Set metadata on a network
	*/
	Label []string

	/*
		Set driver specific options
	*/
	Opt map[string]string

	/*
		Control the network's scope
	*/
	Scope *string

	/*
		Subnet in CIDR format that represents a network segment
	*/
	Subnet *string
}

/*
DockerNetworkCreateCmd is wrapper of 'docker network create'
------------------------------
create [OPTIONS] NETWORK
Create a network
------------------------------
*/
func DockerNetworkCreateCmd(opt DockerNetworkCreateOption, args []string) *exec.Cmd {
	cargs := []string{"network", "create"}
	if opt.Attachable != nil {
		cargs = append(cargs, "--attachable="+fmt.Sprint(*opt.Attachable))
	}
	if opt.AuxAddress != nil {
		for key, val := range opt.AuxAddress {
			cargs = append(cargs, "--aux-address")
			cargs = append(cargs, key+"="+val)
		}
	}
	if opt.ConfigFrom != nil {
		cargs = append(cargs, "--config-from="+fmt.Sprint(*opt.ConfigFrom))
	}
	if opt.ConfigOnly != nil {
		cargs = append(cargs, "--config-only="+fmt.Sprint(*opt.ConfigOnly))
	}
	if opt.Driver != nil {
		cargs = append(cargs, "--driver="+fmt.Sprint(*opt.Driver))
	}
	if opt.Gateway != nil {
		cargs = append(cargs, "--gateway="+fmt.Sprint(*opt.Gateway))
	}
	if opt.Ingress != nil {
		cargs = append(cargs, "--ingress="+fmt.Sprint(*opt.Ingress))
	}
	if opt.Internal != nil {
		cargs = append(cargs, "--internal="+fmt.Sprint(*opt.Internal))
	}
	if opt.IpRange != nil {
		cargs = append(cargs, "--ip-range="+fmt.Sprint(*opt.IpRange))
	}
	if opt.IpamDriver != nil {
		cargs = append(cargs, "--ipam-driver="+fmt.Sprint(*opt.IpamDriver))
	}
	if opt.IpamOpt != nil {
		for key, val := range opt.IpamOpt {
			cargs = append(cargs, "--ipam-opt")
			cargs = append(cargs, key+"="+val)
		}
	}
	if opt.Ipv6 != nil {
		cargs = append(cargs, "--ipv6="+fmt.Sprint(*opt.Ipv6))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.Opt != nil {
		for key, val := range opt.Opt {
			cargs = append(cargs, "--opt")
			cargs = append(cargs, key+"="+val)
		}
	}
	if opt.Scope != nil {
		cargs = append(cargs, "--scope="+fmt.Sprint(*opt.Scope))
	}
	if opt.Subnet != nil {
		cargs = append(cargs, "--subnet="+fmt.Sprint(*opt.Subnet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNetworkDisconnectOption struct {
	/*
		Force the container to disconnect from a network
	*/
	Force *bool
}

/*
DockerNetworkDisconnectCmd is wrapper of 'docker network disconnect'
------------------------------
disconnect [OPTIONS] NETWORK CONTAINER
Disconnect a container from a network
------------------------------
*/
func DockerNetworkDisconnectCmd(opt DockerNetworkDisconnectOption, args []string) *exec.Cmd {
	cargs := []string{"network", "disconnect"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNetworkInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Verbose output for diagnostics
	*/
	Verbose *bool
}

/*
DockerNetworkInspectCmd is wrapper of 'docker network inspect'
------------------------------
inspect [OPTIONS] NETWORK [NETWORK...]
Display detailed information on one or more networks
------------------------------
*/
func DockerNetworkInspectCmd(opt DockerNetworkInspectOption, args []string) *exec.Cmd {
	cargs := []string{"network", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Verbose != nil {
		cargs = append(cargs, "--verbose="+fmt.Sprint(*opt.Verbose))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNetworkLsOption struct {
	/*
		Provide filter values (e.g. 'driver=bridge')
	*/
	Filter *string

	/*
		Pretty-print networks using a Go template
	*/
	Format *string

	/*
		Do not truncate the output
	*/
	NoTrunc *bool

	/*
		Only display network IDs
	*/
	Quiet *bool
}

/*
DockerNetworkLsCmd is wrapper of 'docker network ls'
------------------------------
ls [OPTIONS]
List networks
------------------------------
*/
func DockerNetworkLsCmd(opt DockerNetworkLsOption, args []string) *exec.Cmd {
	cargs := []string{"network", "ls"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNetworkPruneOption struct {
	/*
		Provide filter values (e.g. 'until=<timestamp>')
	*/
	Filter *string

	/*
		Do not prompt for confirmation
	*/
	Force *bool
}

/*
DockerNetworkPruneCmd is wrapper of 'docker network prune'
------------------------------
prune [OPTIONS]
Remove all unused networks
------------------------------
*/
func DockerNetworkPruneCmd(opt DockerNetworkPruneOption, args []string) *exec.Cmd {
	cargs := []string{"network", "prune"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerNetworkRmCmd is wrapper of 'docker network rm'
------------------------------
rm NETWORK [NETWORK...]
Remove one or more networks
------------------------------
*/
func DockerNetworkRmCmd(args []string) *exec.Cmd {
	cargs := []string{"network", "rm"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerNodeCmd is wrapper of 'docker node'
------------------------------
node
Manage Swarm nodes
------------------------------
*/
func DockerNodeCmd(args []string) *exec.Cmd {
	cargs := []string{"node"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerNodeDemoteCmd is wrapper of 'docker node demote'
------------------------------
demote NODE [NODE...]
Demote one or more nodes from manager in the swarm
------------------------------
*/
func DockerNodeDemoteCmd(args []string) *exec.Cmd {
	cargs := []string{"node", "demote"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNodeInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Print the information in a human friendly format
	*/
	Pretty *bool
}

/*
DockerNodeInspectCmd is wrapper of 'docker node inspect'
------------------------------
inspect [OPTIONS] self|NODE [NODE...]
Display detailed information on one or more nodes
------------------------------
*/
func DockerNodeInspectCmd(opt DockerNodeInspectOption, args []string) *exec.Cmd {
	cargs := []string{"node", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Pretty != nil {
		cargs = append(cargs, "--pretty="+fmt.Sprint(*opt.Pretty))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNodeLsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print nodes using a Go template
	*/
	Format *string

	/*
		Only display IDs
	*/
	Quiet *bool
}

/*
DockerNodeLsCmd is wrapper of 'docker node ls'
------------------------------
ls [OPTIONS]
List nodes in the swarm
------------------------------
*/
func DockerNodeLsCmd(opt DockerNodeLsOption, args []string) *exec.Cmd {
	cargs := []string{"node", "ls"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerNodePromoteCmd is wrapper of 'docker node promote'
------------------------------
promote NODE [NODE...]
Promote one or more nodes to manager in the swarm
------------------------------
*/
func DockerNodePromoteCmd(args []string) *exec.Cmd {
	cargs := []string{"node", "promote"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNodePsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print tasks using a Go template
	*/
	Format *string

	/*
		Do not map IDs to Names
	*/
	NoResolve *bool

	/*
		Do not truncate output
	*/
	NoTrunc *bool

	/*
		Only display task IDs
	*/
	Quiet *bool
}

/*
DockerNodePsCmd is wrapper of 'docker node ps'
------------------------------
ps [OPTIONS] [NODE...]
List tasks running on one or more nodes, defaults to current node
------------------------------
*/
func DockerNodePsCmd(opt DockerNodePsOption, args []string) *exec.Cmd {
	cargs := []string{"node", "ps"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoResolve != nil {
		cargs = append(cargs, "--no-resolve="+fmt.Sprint(*opt.NoResolve))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNodeRmOption struct {
	/*
		Force remove a node from the swarm
	*/
	Force *bool
}

/*
DockerNodeRmCmd is wrapper of 'docker node rm'
------------------------------
rm [OPTIONS] NODE [NODE...]
Remove one or more nodes from the swarm
------------------------------
*/
func DockerNodeRmCmd(opt DockerNodeRmOption, args []string) *exec.Cmd {
	cargs := []string{"node", "rm"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerNodeUpdateOption struct {
	/*
		Availability of the node ("active"|"pause"|"drain")
	*/
	Availability *string

	/*
		Add or update a node label (key=value)
	*/
	LabelAdd []string

	/*
		Remove a node label if exists
	*/
	LabelRm []string

	/*
		Role of the node ("worker"|"manager")
	*/
	Role *string
}

/*
DockerNodeUpdateCmd is wrapper of 'docker node update'
------------------------------
update [OPTIONS] NODE
Update a node
------------------------------
*/
func DockerNodeUpdateCmd(opt DockerNodeUpdateOption, args []string) *exec.Cmd {
	cargs := []string{"node", "update"}
	if opt.Availability != nil {
		cargs = append(cargs, "--availability="+fmt.Sprint(*opt.Availability))
	}
	if opt.LabelAdd != nil {
		for _, str := range opt.LabelAdd {
			cargs = append(cargs, "--label-add")
			cargs = append(cargs, str)
		}
	}
	if opt.LabelRm != nil {
		for _, str := range opt.LabelRm {
			cargs = append(cargs, "--label-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.Role != nil {
		cargs = append(cargs, "--role="+fmt.Sprint(*opt.Role))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerPauseCmd is wrapper of 'docker pause'
------------------------------
pause CONTAINER [CONTAINER...]
Pause all processes within one or more containers
------------------------------
*/
func DockerPauseCmd(args []string) *exec.Cmd {
	cargs := []string{"pause"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerPluginCmd is wrapper of 'docker plugin'
------------------------------
plugin
Manage plugins
------------------------------
*/
func DockerPluginCmd(args []string) *exec.Cmd {
	cargs := []string{"plugin"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginCreateOption struct {
	/*
		Compress the context using gzip
	*/
	Compress *bool
}

/*
DockerPluginCreateCmd is wrapper of 'docker plugin create'
------------------------------
create [OPTIONS] PLUGIN PLUGIN-DATA-DIR
Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory.
------------------------------
*/
func DockerPluginCreateCmd(opt DockerPluginCreateOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "create"}
	if opt.Compress != nil {
		cargs = append(cargs, "--compress="+fmt.Sprint(*opt.Compress))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginDisableOption struct {
	/*
		Force the disable of an active plugin
	*/
	Force *bool
}

/*
DockerPluginDisableCmd is wrapper of 'docker plugin disable'
------------------------------
disable [OPTIONS] PLUGIN
Disable a plugin
------------------------------
*/
func DockerPluginDisableCmd(opt DockerPluginDisableOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "disable"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginEnableOption struct {
	/*
		HTTP client timeout (in seconds)
	*/
	Timeout *int
}

/*
DockerPluginEnableCmd is wrapper of 'docker plugin enable'
------------------------------
enable [OPTIONS] PLUGIN
Enable a plugin
------------------------------
*/
func DockerPluginEnableCmd(opt DockerPluginEnableOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "enable"}
	if opt.Timeout != nil {
		cargs = append(cargs, "--timeout="+fmt.Sprint(*opt.Timeout))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string
}

/*
DockerPluginInspectCmd is wrapper of 'docker plugin inspect'
------------------------------
inspect [OPTIONS] PLUGIN [PLUGIN...]
Display detailed information on one or more plugins
------------------------------
*/
func DockerPluginInspectCmd(opt DockerPluginInspectOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginInstallOption struct {
	/*
		Local name for plugin
	*/
	Alias *string

	/*
		Do not enable the plugin on install
	*/
	Disable *bool

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Grant all permissions necessary to run the plugin
	*/
	GrantAllPermissions *bool
}

/*
DockerPluginInstallCmd is wrapper of 'docker plugin install'
------------------------------
install [OPTIONS] PLUGIN [KEY=VALUE...]
Install a plugin
------------------------------
*/
func DockerPluginInstallCmd(opt DockerPluginInstallOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "install"}
	if opt.Alias != nil {
		cargs = append(cargs, "--alias="+fmt.Sprint(*opt.Alias))
	}
	if opt.Disable != nil {
		cargs = append(cargs, "--disable="+fmt.Sprint(*opt.Disable))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.GrantAllPermissions != nil {
		cargs = append(cargs, "--grant-all-permissions="+fmt.Sprint(*opt.GrantAllPermissions))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginLsOption struct {
	/*
		Provide filter values (e.g. 'enabled=true')
	*/
	Filter *string

	/*
		Pretty-print plugins using a Go template
	*/
	Format *string

	/*
		Don't truncate output
	*/
	NoTrunc *bool

	/*
		Only display plugin IDs
	*/
	Quiet *bool
}

/*
DockerPluginLsCmd is wrapper of 'docker plugin ls'
------------------------------
ls [OPTIONS]
List plugins
------------------------------
*/
func DockerPluginLsCmd(opt DockerPluginLsOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "ls"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginPushOption struct {
	/*
		Skip image signing
	*/
	DisableContentTrust *bool
}

/*
DockerPluginPushCmd is wrapper of 'docker plugin push'
------------------------------
push [OPTIONS] PLUGIN[:TAG]
Push a plugin to a registry
------------------------------
*/
func DockerPluginPushCmd(opt DockerPluginPushOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "push"}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginRmOption struct {
	/*
		Force the removal of an active plugin
	*/
	Force *bool
}

/*
DockerPluginRmCmd is wrapper of 'docker plugin rm'
------------------------------
rm [OPTIONS] PLUGIN [PLUGIN...]
Remove one or more plugins
------------------------------
*/
func DockerPluginRmCmd(opt DockerPluginRmOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "rm"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerPluginSetCmd is wrapper of 'docker plugin set'
------------------------------
set PLUGIN KEY=VALUE [KEY=VALUE...]
Change settings for a plugin
------------------------------
*/
func DockerPluginSetCmd(args []string) *exec.Cmd {
	cargs := []string{"plugin", "set"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPluginUpgradeOption struct {
	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Grant all permissions necessary to run the plugin
	*/
	GrantAllPermissions *bool

	/*
		Do not check if specified remote plugin matches existing plugin image
	*/
	SkipRemoteCheck *bool
}

/*
DockerPluginUpgradeCmd is wrapper of 'docker plugin upgrade'
------------------------------
upgrade [OPTIONS] PLUGIN [REMOTE]
Upgrade an existing plugin
------------------------------
*/
func DockerPluginUpgradeCmd(opt DockerPluginUpgradeOption, args []string) *exec.Cmd {
	cargs := []string{"plugin", "upgrade"}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.GrantAllPermissions != nil {
		cargs = append(cargs, "--grant-all-permissions="+fmt.Sprint(*opt.GrantAllPermissions))
	}
	if opt.SkipRemoteCheck != nil {
		cargs = append(cargs, "--skip-remote-check="+fmt.Sprint(*opt.SkipRemoteCheck))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerPortCmd is wrapper of 'docker port'
------------------------------
port CONTAINER [PRIVATE_PORT[/PROTO]]
List port mappings or a specific mapping for the container
------------------------------
*/
func DockerPortCmd(args []string) *exec.Cmd {
	cargs := []string{"port"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPsOption struct {
	/*
		Show all containers (default shows just running)
	*/
	All *bool

	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print containers using a Go template
	*/
	Format *string

	/*
		Show n last created containers (includes all states)
	*/
	Last *int

	/*
		Show the latest created container (includes all states)
	*/
	Latest *bool

	/*
		Don't truncate output
	*/
	NoTrunc *bool

	/*
		Only display container IDs
	*/
	Quiet *bool

	/*
		Display total file sizes
	*/
	Size *bool
}

/*
DockerPsCmd is wrapper of 'docker ps'
------------------------------
ps [OPTIONS]
List containers
------------------------------
*/
func DockerPsCmd(opt DockerPsOption, args []string) *exec.Cmd {
	cargs := []string{"ps"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Last != nil {
		cargs = append(cargs, "--last="+fmt.Sprint(*opt.Last))
	}
	if opt.Latest != nil {
		cargs = append(cargs, "--latest="+fmt.Sprint(*opt.Latest))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Size != nil {
		cargs = append(cargs, "--size="+fmt.Sprint(*opt.Size))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPullOption struct {
	/*
		Download all tagged images in the repository
	*/
	AllTags *bool

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Suppress verbose output
	*/
	Quiet *bool
}

/*
DockerPullCmd is wrapper of 'docker pull'
------------------------------
pull [OPTIONS] NAME[:TAG|@DIGEST]
Pull an image or a repository from a registry
------------------------------
*/
func DockerPullCmd(opt DockerPullOption, args []string) *exec.Cmd {
	cargs := []string{"pull"}
	if opt.AllTags != nil {
		cargs = append(cargs, "--all-tags="+fmt.Sprint(*opt.AllTags))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerPushOption struct {
	/*
		Push all tagged images in the repository
	*/
	AllTags *bool

	/*
		Skip image signing
	*/
	DisableContentTrust *bool

	/*
		Suppress verbose output
	*/
	Quiet *bool
}

/*
DockerPushCmd is wrapper of 'docker push'
------------------------------
push [OPTIONS] NAME[:TAG]
Push an image or a repository to a registry
------------------------------
*/
func DockerPushCmd(opt DockerPushOption, args []string) *exec.Cmd {
	cargs := []string{"push"}
	if opt.AllTags != nil {
		cargs = append(cargs, "--all-tags="+fmt.Sprint(*opt.AllTags))
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerRenameCmd is wrapper of 'docker rename'
------------------------------
rename CONTAINER NEW_NAME
Rename a container
------------------------------
*/
func DockerRenameCmd(args []string) *exec.Cmd {
	cargs := []string{"rename"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerRestartOption struct {
	/*
		Seconds to wait for stop before killing the container
	*/
	Time *int
}

/*
DockerRestartCmd is wrapper of 'docker restart'
------------------------------
restart [OPTIONS] CONTAINER [CONTAINER...]
Restart one or more containers
------------------------------
*/
func DockerRestartCmd(opt DockerRestartOption, args []string) *exec.Cmd {
	cargs := []string{"restart"}
	if opt.Time != nil {
		cargs = append(cargs, "--time="+fmt.Sprint(*opt.Time))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerRmOption struct {
	/*
		Force the removal of a running container (uses SIGKILL)
	*/
	Force *bool

	/*
		Remove the specified link
	*/
	Link *bool

	/*
		Remove anonymous volumes associated with the container
	*/
	Volumes *bool
}

/*
DockerRmCmd is wrapper of 'docker rm'
------------------------------
rm [OPTIONS] CONTAINER [CONTAINER...]
Remove one or more containers
------------------------------
*/
func DockerRmCmd(opt DockerRmOption, args []string) *exec.Cmd {
	cargs := []string{"rm"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	if opt.Link != nil {
		cargs = append(cargs, "--link="+fmt.Sprint(*opt.Link))
	}
	if opt.Volumes != nil {
		cargs = append(cargs, "--volumes="+fmt.Sprint(*opt.Volumes))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerRmiOption struct {
	/*
		Force removal of the image
	*/
	Force *bool

	/*
		Do not delete untagged parents
	*/
	NoPrune *bool
}

/*
DockerRmiCmd is wrapper of 'docker rmi'
------------------------------
rmi [OPTIONS] IMAGE [IMAGE...]
Remove one or more images
------------------------------
*/
func DockerRmiCmd(opt DockerRmiOption, args []string) *exec.Cmd {
	cargs := []string{"rmi"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	if opt.NoPrune != nil {
		cargs = append(cargs, "--no-prune="+fmt.Sprint(*opt.NoPrune))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerRunOption struct {
	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	AddHost []string

	/*
		Attach to STDIN, STDOUT or STDERR
	*/
	Attach []string

	/*
		Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
	*/
	BlkioWeight *uint16

	/*
		Block IO weight (relative device weight)
	*/
	BlkioWeightDevice []string

	/*
		Add Linux capabilities
	*/
	CapAdd []string

	/*
		Drop Linux capabilities
	*/
	CapDrop []string

	/*
		Optional parent cgroup for the container
	*/
	CgroupParent *string

	/*
			Cgroup namespace to use (host|private)
		'host':    Run the container in the Docker host's cgroup namespace
		'private': Run the container in its own private cgroup namespace
		'':        Use the cgroup namespace as configured by the
		           default-cgroupns-mode option on the daemon (default)
	*/
	Cgroupns *string

	/*
		Write the container ID to the file
	*/
	Cidfile *string

	/*
		CPU count (Windows only)
	*/
	CpuCount *int64

	/*
		CPU percent (Windows only)
	*/
	CpuPercent *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		Limit CPU real-time period in microseconds
	*/
	CpuRtPeriod *int64

	/*
		Limit CPU real-time runtime in microseconds
	*/
	CpuRtRuntime *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		Number of CPUs
	*/
	Cpus *string

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Run container in background and print container ID
	*/
	Detach *bool

	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Add a host device to the container
	*/
	Device []string

	/*
		Add a rule to the cgroup allowed devices list
	*/
	DeviceCgroupRule []string

	/*
		Limit read rate (bytes per second) from a device
	*/
	DeviceReadBps []string

	/*
		Limit read rate (IO per second) from a device
	*/
	DeviceReadIops []string

	/*
		Limit write rate (bytes per second) to a device
	*/
	DeviceWriteBps []string

	/*
		Limit write rate (IO per second) to a device
	*/
	DeviceWriteIops []string

	/*
		Skip image verification
	*/
	DisableContentTrust *bool

	/*
		Set custom DNS servers
	*/
	Dns []string

	/*
		Set DNS options
	*/
	DnsOpt []string

	/*
		Set DNS options
	*/
	DnsOption []string

	/*
		Set custom DNS search domains
	*/
	DnsSearch []string

	/*
		Container NIS domain name
	*/
	Domainname *string

	/*
		Overwrite the default ENTRYPOINT of the image
	*/
	Entrypoint *string

	/*
		Set environment variables
	*/
	Env []string

	/*
		Read in a file of environment variables
	*/
	EnvFile []string

	/*
		Expose a port or a range of ports
	*/
	Expose []string

	/*
		GPU devices to add to the container ('all' to pass all GPUs)
	*/
	Gpus *string

	/*
		Add additional groups to join
	*/
	GroupAdd []string

	/*
		Command to run to check health
	*/
	HealthCmd *string

	/*
		Time between running the check (ms|s|m|h) (default 0s)
	*/
	HealthInterval *string

	/*
		Consecutive failures needed to report unhealthy
	*/
	HealthRetries *int

	/*
		Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
	*/
	HealthStartPeriod *string

	/*
		Maximum time to allow one check to run (ms|s|m|h) (default 0s)
	*/
	HealthTimeout *string

	/*
		Print usage
	*/
	Help *bool

	/*
		Container host name
	*/
	Hostname *string

	/*
		Run an init inside the container that forwards signals and reaps processes
	*/
	Init *bool

	/*
		Keep STDIN open even if not attached
	*/
	Interactive *bool

	/*
		Maximum IO bandwidth limit for the system drive (Windows only)
	*/
	IoMaxbandwidth *string

	/*
		Maximum IOps limit for the system drive (Windows only)
	*/
	IoMaxiops *uint64

	/*
		IPv4 address (e.g., 172.30.100.104)
	*/
	Ip *string

	/*
		IPv6 address (e.g., 2001:db8::33)
	*/
	Ip6 *string

	/*
		IPC mode to use
	*/
	Ipc *string

	/*
		Container isolation technology
	*/
	Isolation *string

	/*
		Kernel memory limit
	*/
	KernelMemory *string

	/*
		Set meta data on a container
	*/
	Label []string

	/*
		Read in a line delimited file of labels
	*/
	LabelFile []string

	/*
		Add link to another container
	*/
	Link []string

	/*
		Container IPv4/IPv6 link-local addresses
	*/
	LinkLocalIp []string

	/*
		Logging driver for the container
	*/
	LogDriver *string

	/*
		Log driver options
	*/
	LogOpt []string

	/*
		Container MAC address (e.g., 92:d0:c6:0a:29:33)
	*/
	MacAddress *string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Memory soft limit
	*/
	MemoryReservation *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Tune container memory swappiness (0 to 100)
	*/
	MemorySwappiness *int64

	/*
		Attach a filesystem mount to the container
	*/
	Mount *string

	/*
		Assign a name to the container
	*/
	Name *string

	/*
		Connect a container to a network
	*/
	Net *string

	/*
		Add network-scoped alias for the container
	*/
	NetAlias []string

	/*
		Connect a container to a network
	*/
	Network *string

	/*
		Add network-scoped alias for the container
	*/
	NetworkAlias []string

	/*
		Disable any container-specified HEALTHCHECK
	*/
	NoHealthcheck *bool

	/*
		Disable OOM Killer
	*/
	OomKillDisable *bool

	/*
		Tune host's OOM preferences (-1000 to 1000)
	*/
	OomScoreAdj *int

	/*
		PID namespace to use
	*/
	Pid *string

	/*
		Tune container pids limit (set -1 for unlimited)
	*/
	PidsLimit *int64

	/*
		Set platform if server is multi-platform capable
	*/
	Platform *string

	/*
		Give extended privileges to this container
	*/
	Privileged *bool

	/*
		Publish a container's port(s) to the host
	*/
	Publish []string

	/*
		Publish all exposed ports to random ports
	*/
	PublishAll *bool

	/*
		Pull image before running ("always"|"missing"|"never")
	*/
	Pull *string

	/*
		Mount the container's root filesystem as read only
	*/
	ReadOnly *bool

	/*
		Restart policy to apply when a container exits
	*/
	Restart *string

	/*
		Automatically remove the container when it exits
	*/
	Rm *bool

	/*
		Runtime to use for this container
	*/
	Runtime *string

	/*
		Security Options
	*/
	SecurityOpt []string

	/*
		Size of /dev/shm
	*/
	ShmSize *string

	/*
		Proxy received signals to the process
	*/
	SigProxy *bool

	/*
		Signal to stop a container
	*/
	StopSignal *string

	/*
		Timeout (in seconds) to stop a container
	*/
	StopTimeout *int

	/*
		Storage driver options for the container
	*/
	StorageOpt []string

	/*
		Sysctl options
	*/
	Sysctl map[string]string

	/*
		Mount a tmpfs directory
	*/
	Tmpfs []string

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Ulimit options
	*/
	Ulimit *string

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		User namespace to use
	*/
	Userns *string

	/*
		UTS namespace to use
	*/
	Uts *string

	/*
		Bind mount a volume
	*/
	Volume []string

	/*
		Optional volume driver for the container
	*/
	VolumeDriver *string

	/*
		Mount volumes from the specified container(s)
	*/
	VolumesFrom []string

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerRunCmd is wrapper of 'docker run'
------------------------------
run [OPTIONS] IMAGE [COMMAND] [ARG...]
Run a command in a new container
------------------------------
*/
func DockerRunCmd(opt DockerRunOption, args []string) *exec.Cmd {
	cargs := []string{"run"}
	if opt.AddHost != nil {
		for _, str := range opt.AddHost {
			cargs = append(cargs, "--add-host")
			cargs = append(cargs, str)
		}
	}
	if opt.Attach != nil {
		for _, str := range opt.Attach {
			cargs = append(cargs, "--attach")
			cargs = append(cargs, str)
		}
	}
	if opt.BlkioWeight != nil {
		cargs = append(cargs, "--blkio-weight="+fmt.Sprint(*opt.BlkioWeight))
	}
	if opt.BlkioWeightDevice != nil {
		for _, str := range opt.BlkioWeightDevice {
			cargs = append(cargs, "--blkio-weight-device")
			cargs = append(cargs, str)
		}
	}
	if opt.CapAdd != nil {
		for _, str := range opt.CapAdd {
			cargs = append(cargs, "--cap-add")
			cargs = append(cargs, str)
		}
	}
	if opt.CapDrop != nil {
		for _, str := range opt.CapDrop {
			cargs = append(cargs, "--cap-drop")
			cargs = append(cargs, str)
		}
	}
	if opt.CgroupParent != nil {
		cargs = append(cargs, "--cgroup-parent="+fmt.Sprint(*opt.CgroupParent))
	}
	if opt.Cgroupns != nil {
		cargs = append(cargs, "--cgroupns="+fmt.Sprint(*opt.Cgroupns))
	}
	if opt.Cidfile != nil {
		cargs = append(cargs, "--cidfile="+fmt.Sprint(*opt.Cidfile))
	}
	if opt.CpuCount != nil {
		cargs = append(cargs, "--cpu-count="+fmt.Sprint(*opt.CpuCount))
	}
	if opt.CpuPercent != nil {
		cargs = append(cargs, "--cpu-percent="+fmt.Sprint(*opt.CpuPercent))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuRtPeriod != nil {
		cargs = append(cargs, "--cpu-rt-period="+fmt.Sprint(*opt.CpuRtPeriod))
	}
	if opt.CpuRtRuntime != nil {
		cargs = append(cargs, "--cpu-rt-runtime="+fmt.Sprint(*opt.CpuRtRuntime))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.Cpus != nil {
		cargs = append(cargs, "--cpus="+fmt.Sprint(*opt.Cpus))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.Device != nil {
		for _, str := range opt.Device {
			cargs = append(cargs, "--device")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceCgroupRule != nil {
		for _, str := range opt.DeviceCgroupRule {
			cargs = append(cargs, "--device-cgroup-rule")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadBps != nil {
		for _, str := range opt.DeviceReadBps {
			cargs = append(cargs, "--device-read-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceReadIops != nil {
		for _, str := range opt.DeviceReadIops {
			cargs = append(cargs, "--device-read-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteBps != nil {
		for _, str := range opt.DeviceWriteBps {
			cargs = append(cargs, "--device-write-bps")
			cargs = append(cargs, str)
		}
	}
	if opt.DeviceWriteIops != nil {
		for _, str := range opt.DeviceWriteIops {
			cargs = append(cargs, "--device-write-iops")
			cargs = append(cargs, str)
		}
	}
	if opt.DisableContentTrust != nil {
		cargs = append(cargs, "--disable-content-trust="+fmt.Sprint(*opt.DisableContentTrust))
	}
	if opt.Dns != nil {
		for _, str := range opt.Dns {
			cargs = append(cargs, "--dns")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOpt != nil {
		for _, str := range opt.DnsOpt {
			cargs = append(cargs, "--dns-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOption != nil {
		for _, str := range opt.DnsOption {
			cargs = append(cargs, "--dns-option")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsSearch != nil {
		for _, str := range opt.DnsSearch {
			cargs = append(cargs, "--dns-search")
			cargs = append(cargs, str)
		}
	}
	if opt.Domainname != nil {
		cargs = append(cargs, "--domainname="+fmt.Sprint(*opt.Domainname))
	}
	if opt.Entrypoint != nil {
		cargs = append(cargs, "--entrypoint="+fmt.Sprint(*opt.Entrypoint))
	}
	if opt.Env != nil {
		for _, str := range opt.Env {
			cargs = append(cargs, "--env")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvFile != nil {
		for _, str := range opt.EnvFile {
			cargs = append(cargs, "--env-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Expose != nil {
		for _, str := range opt.Expose {
			cargs = append(cargs, "--expose")
			cargs = append(cargs, str)
		}
	}
	if opt.Gpus != nil {
		cargs = append(cargs, "--gpus="+fmt.Sprint(*opt.Gpus))
	}
	if opt.GroupAdd != nil {
		for _, str := range opt.GroupAdd {
			cargs = append(cargs, "--group-add")
			cargs = append(cargs, str)
		}
	}
	if opt.HealthCmd != nil {
		cargs = append(cargs, "--health-cmd="+fmt.Sprint(*opt.HealthCmd))
	}
	if opt.HealthInterval != nil {
		cargs = append(cargs, "--health-interval="+fmt.Sprint(*opt.HealthInterval))
	}
	if opt.HealthRetries != nil {
		cargs = append(cargs, "--health-retries="+fmt.Sprint(*opt.HealthRetries))
	}
	if opt.HealthStartPeriod != nil {
		cargs = append(cargs, "--health-start-period="+fmt.Sprint(*opt.HealthStartPeriod))
	}
	if opt.HealthTimeout != nil {
		cargs = append(cargs, "--health-timeout="+fmt.Sprint(*opt.HealthTimeout))
	}
	if opt.Help != nil {
		cargs = append(cargs, "--help="+fmt.Sprint(*opt.Help))
	}
	if opt.Hostname != nil {
		cargs = append(cargs, "--hostname="+fmt.Sprint(*opt.Hostname))
	}
	if opt.Init != nil {
		cargs = append(cargs, "--init="+fmt.Sprint(*opt.Init))
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	if opt.IoMaxbandwidth != nil {
		cargs = append(cargs, "--io-maxbandwidth="+fmt.Sprint(*opt.IoMaxbandwidth))
	}
	if opt.IoMaxiops != nil {
		cargs = append(cargs, "--io-maxiops="+fmt.Sprint(*opt.IoMaxiops))
	}
	if opt.Ip != nil {
		cargs = append(cargs, "--ip="+fmt.Sprint(*opt.Ip))
	}
	if opt.Ip6 != nil {
		cargs = append(cargs, "--ip6="+fmt.Sprint(*opt.Ip6))
	}
	if opt.Ipc != nil {
		cargs = append(cargs, "--ipc="+fmt.Sprint(*opt.Ipc))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.KernelMemory != nil {
		cargs = append(cargs, "--kernel-memory="+fmt.Sprint(*opt.KernelMemory))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.LabelFile != nil {
		for _, str := range opt.LabelFile {
			cargs = append(cargs, "--label-file")
			cargs = append(cargs, str)
		}
	}
	if opt.Link != nil {
		for _, str := range opt.Link {
			cargs = append(cargs, "--link")
			cargs = append(cargs, str)
		}
	}
	if opt.LinkLocalIp != nil {
		for _, str := range opt.LinkLocalIp {
			cargs = append(cargs, "--link-local-ip")
			cargs = append(cargs, str)
		}
	}
	if opt.LogDriver != nil {
		cargs = append(cargs, "--log-driver="+fmt.Sprint(*opt.LogDriver))
	}
	if opt.LogOpt != nil {
		for _, str := range opt.LogOpt {
			cargs = append(cargs, "--log-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.MacAddress != nil {
		cargs = append(cargs, "--mac-address="+fmt.Sprint(*opt.MacAddress))
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemoryReservation != nil {
		cargs = append(cargs, "--memory-reservation="+fmt.Sprint(*opt.MemoryReservation))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.MemorySwappiness != nil {
		cargs = append(cargs, "--memory-swappiness="+fmt.Sprint(*opt.MemorySwappiness))
	}
	if opt.Mount != nil {
		cargs = append(cargs, "--mount="+fmt.Sprint(*opt.Mount))
	}
	if opt.Name != nil {
		cargs = append(cargs, "--name="+fmt.Sprint(*opt.Name))
	}
	if opt.Net != nil {
		cargs = append(cargs, "--net="+fmt.Sprint(*opt.Net))
	}
	if opt.NetAlias != nil {
		for _, str := range opt.NetAlias {
			cargs = append(cargs, "--net-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NetworkAlias != nil {
		for _, str := range opt.NetworkAlias {
			cargs = append(cargs, "--network-alias")
			cargs = append(cargs, str)
		}
	}
	if opt.NoHealthcheck != nil {
		cargs = append(cargs, "--no-healthcheck="+fmt.Sprint(*opt.NoHealthcheck))
	}
	if opt.OomKillDisable != nil {
		cargs = append(cargs, "--oom-kill-disable="+fmt.Sprint(*opt.OomKillDisable))
	}
	if opt.OomScoreAdj != nil {
		cargs = append(cargs, "--oom-score-adj="+fmt.Sprint(*opt.OomScoreAdj))
	}
	if opt.Pid != nil {
		cargs = append(cargs, "--pid="+fmt.Sprint(*opt.Pid))
	}
	if opt.PidsLimit != nil {
		cargs = append(cargs, "--pids-limit="+fmt.Sprint(*opt.PidsLimit))
	}
	if opt.Platform != nil {
		cargs = append(cargs, "--platform="+fmt.Sprint(*opt.Platform))
	}
	if opt.Privileged != nil {
		cargs = append(cargs, "--privileged="+fmt.Sprint(*opt.Privileged))
	}
	if opt.Publish != nil {
		for _, str := range opt.Publish {
			cargs = append(cargs, "--publish")
			cargs = append(cargs, str)
		}
	}
	if opt.PublishAll != nil {
		cargs = append(cargs, "--publish-all="+fmt.Sprint(*opt.PublishAll))
	}
	if opt.Pull != nil {
		cargs = append(cargs, "--pull="+fmt.Sprint(*opt.Pull))
	}
	if opt.ReadOnly != nil {
		cargs = append(cargs, "--read-only="+fmt.Sprint(*opt.ReadOnly))
	}
	if opt.Restart != nil {
		cargs = append(cargs, "--restart="+fmt.Sprint(*opt.Restart))
	}
	if opt.Rm != nil {
		cargs = append(cargs, "--rm="+fmt.Sprint(*opt.Rm))
	}
	if opt.Runtime != nil {
		cargs = append(cargs, "--runtime="+fmt.Sprint(*opt.Runtime))
	}
	if opt.SecurityOpt != nil {
		for _, str := range opt.SecurityOpt {
			cargs = append(cargs, "--security-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.ShmSize != nil {
		cargs = append(cargs, "--shm-size="+fmt.Sprint(*opt.ShmSize))
	}
	if opt.SigProxy != nil {
		cargs = append(cargs, "--sig-proxy="+fmt.Sprint(*opt.SigProxy))
	}
	if opt.StopSignal != nil {
		cargs = append(cargs, "--stop-signal="+fmt.Sprint(*opt.StopSignal))
	}
	if opt.StopTimeout != nil {
		cargs = append(cargs, "--stop-timeout="+fmt.Sprint(*opt.StopTimeout))
	}
	if opt.StorageOpt != nil {
		for _, str := range opt.StorageOpt {
			cargs = append(cargs, "--storage-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.Sysctl != nil {
		for key, val := range opt.Sysctl {
			cargs = append(cargs, "--sysctl")
			cargs = append(cargs, key+"="+val)
		}
	}
	if opt.Tmpfs != nil {
		for _, str := range opt.Tmpfs {
			cargs = append(cargs, "--tmpfs")
			cargs = append(cargs, str)
		}
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.Userns != nil {
		cargs = append(cargs, "--userns="+fmt.Sprint(*opt.Userns))
	}
	if opt.Uts != nil {
		cargs = append(cargs, "--uts="+fmt.Sprint(*opt.Uts))
	}
	if opt.Volume != nil {
		for _, str := range opt.Volume {
			cargs = append(cargs, "--volume")
			cargs = append(cargs, str)
		}
	}
	if opt.VolumeDriver != nil {
		cargs = append(cargs, "--volume-driver="+fmt.Sprint(*opt.VolumeDriver))
	}
	if opt.VolumesFrom != nil {
		for _, str := range opt.VolumesFrom {
			cargs = append(cargs, "--volumes-from")
			cargs = append(cargs, str)
		}
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSaveOption struct {
	/*
		Write to a file, instead of STDOUT
	*/
	Output *string
}

/*
DockerSaveCmd is wrapper of 'docker save'
------------------------------
save [OPTIONS] IMAGE [IMAGE...]
Save one or more images to a tar archive (streamed to STDOUT by default)
------------------------------
*/
func DockerSaveCmd(opt DockerSaveOption, args []string) *exec.Cmd {
	cargs := []string{"save"}
	if opt.Output != nil {
		cargs = append(cargs, "--output="+fmt.Sprint(*opt.Output))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSearchOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print search using a Go template
	*/
	Format *string

	/*
		Max number of search results
	*/
	Limit *int

	/*
		Don't truncate output
	*/
	NoTrunc *bool
}

/*
DockerSearchCmd is wrapper of 'docker search'
------------------------------
search [OPTIONS] TERM
Search the Docker Hub for images
------------------------------
*/
func DockerSearchCmd(opt DockerSearchOption, args []string) *exec.Cmd {
	cargs := []string{"search"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Limit != nil {
		cargs = append(cargs, "--limit="+fmt.Sprint(*opt.Limit))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerSecretCmd is wrapper of 'docker secret'
------------------------------
secret
Manage Docker secrets
------------------------------
*/
func DockerSecretCmd(args []string) *exec.Cmd {
	cargs := []string{"secret"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSecretCreateOption struct {
	/*
		Secret driver
	*/
	Driver *string

	/*
		Secret labels
	*/
	Label []string

	/*
		Template driver
	*/
	TemplateDriver *string
}

/*
DockerSecretCreateCmd is wrapper of 'docker secret create'
------------------------------
create [OPTIONS] SECRET [file|-]
Create a secret from a file or STDIN as content
------------------------------
*/
func DockerSecretCreateCmd(opt DockerSecretCreateOption, args []string) *exec.Cmd {
	cargs := []string{"secret", "create"}
	if opt.Driver != nil {
		cargs = append(cargs, "--driver="+fmt.Sprint(*opt.Driver))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.TemplateDriver != nil {
		cargs = append(cargs, "--template-driver="+fmt.Sprint(*opt.TemplateDriver))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSecretInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Print the information in a human friendly format
	*/
	Pretty *bool
}

/*
DockerSecretInspectCmd is wrapper of 'docker secret inspect'
------------------------------
inspect [OPTIONS] SECRET [SECRET...]
Display detailed information on one or more secrets
------------------------------
*/
func DockerSecretInspectCmd(opt DockerSecretInspectOption, args []string) *exec.Cmd {
	cargs := []string{"secret", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Pretty != nil {
		cargs = append(cargs, "--pretty="+fmt.Sprint(*opt.Pretty))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSecretLsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print secrets using a Go template
	*/
	Format *string

	/*
		Only display IDs
	*/
	Quiet *bool
}

/*
DockerSecretLsCmd is wrapper of 'docker secret ls'
------------------------------
ls [OPTIONS]
List secrets
------------------------------
*/
func DockerSecretLsCmd(opt DockerSecretLsOption, args []string) *exec.Cmd {
	cargs := []string{"secret", "ls"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerSecretRmCmd is wrapper of 'docker secret rm'
------------------------------
rm SECRET [SECRET...]
Remove one or more secrets
------------------------------
*/
func DockerSecretRmCmd(args []string) *exec.Cmd {
	cargs := []string{"secret", "rm"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerServiceCmd is wrapper of 'docker service'
------------------------------
service
Manage services
------------------------------
*/
func DockerServiceCmd(args []string) *exec.Cmd {
	cargs := []string{"service"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServiceCreateOption struct {
	/*
		Add Linux capabilities
	*/
	CapAdd []string

	/*
		Drop Linux capabilities
	*/
	CapDrop []string

	/*
		Specify configurations to expose to the service
	*/
	Config *string

	/*
		Placement constraints
	*/
	Constraint []string

	/*
		Container labels
	*/
	ContainerLabel []string

	/*
		Credential spec for managed service account (Windows only)
	*/
	CredentialSpec *string

	/*
		Exit immediately instead of waiting for the service to converge
	*/
	Detach *bool

	/*
		Set custom DNS servers
	*/
	Dns []string

	/*
		Set DNS options
	*/
	DnsOption []string

	/*
		Set custom DNS search domains
	*/
	DnsSearch []string

	/*
		Endpoint mode (vip or dnsrr)
	*/
	EndpointMode *string

	/*
		Overwrite the default ENTRYPOINT of the image
	*/
	Entrypoint *string

	/*
		Set environment variables
	*/
	Env []string

	/*
		Read in a file of environment variables
	*/
	EnvFile []string

	/*
		User defined resources
	*/
	GenericResource []string

	/*
		Set one or more supplementary user groups for the container
	*/
	Group []string

	/*
		Command to run to check health
	*/
	HealthCmd *string

	/*
		Time between running the check (ms|s|m|h)
	*/
	HealthInterval *string

	/*
		Consecutive failures needed to report unhealthy
	*/
	HealthRetries *int

	/*
		Start period for the container to initialize before counting retries towards unstable (ms|s|m|h)
	*/
	HealthStartPeriod *string

	/*
		Maximum time to allow one check to run (ms|s|m|h)
	*/
	HealthTimeout *string

	/*
		Set one or more custom host-to-IP mappings (host:ip)
	*/
	Host []string

	/*
		Container hostname
	*/
	Hostname *string

	/*
		Use an init inside each service container to forward signals and reap processes
	*/
	Init *bool

	/*
		Service container isolation mode
	*/
	Isolation *string

	/*
		Service labels
	*/
	Label []string

	/*
		Limit CPUs
	*/
	LimitCpu *string

	/*
		Limit Memory
	*/
	LimitMemory *string

	/*
		Limit maximum number of processes (default 0 = unlimited)
	*/
	LimitPids *int64

	/*
		Logging driver for service
	*/
	LogDriver *string

	/*
		Logging driver options
	*/
	LogOpt []string

	/*
		Number of job tasks to run concurrently (default equal to --replicas)
	*/
	MaxConcurrent *uint

	/*
		Service mode (replicated, global, replicated-job, or global-job)
	*/
	Mode *string

	/*
		Attach a filesystem mount to the service
	*/
	Mount *string

	/*
		Service name
	*/
	Name *string

	/*
		Network attachments
	*/
	Network *string

	/*
		Disable any container-specified HEALTHCHECK
	*/
	NoHealthcheck *bool

	/*
		Do not query the registry to resolve image digest and supported platforms
	*/
	NoResolveImage *bool

	/*
		Add a placement preference
	*/
	PlacementPref *string

	/*
		Publish a port as a node port
	*/
	Publish *string

	/*
		Suppress progress output
	*/
	Quiet *bool

	/*
		Mount the container's root filesystem as read only
	*/
	ReadOnly *bool

	/*
		Number of tasks
	*/
	Replicas *uint

	/*
		Maximum number of tasks per node (default 0 = unlimited)
	*/
	ReplicasMaxPerNode *uint64

	/*
		Reserve CPUs
	*/
	ReserveCpu *string

	/*
		Reserve Memory
	*/
	ReserveMemory *string

	/*
		Restart when condition is met ("none"|"on-failure"|"any") (default "any")
	*/
	RestartCondition *string

	/*
		Delay between restart attempts (ns|us|ms|s|m|h) (default 5s)
	*/
	RestartDelay *string

	/*
		Maximum number of restarts before giving up
	*/
	RestartMaxAttempts *uint

	/*
		Window used to evaluate the restart policy (ns|us|ms|s|m|h)
	*/
	RestartWindow *string

	/*
		Delay between task rollbacks (ns|us|ms|s|m|h) (default 0s)
	*/
	RollbackDelay *string

	/*
		Action on rollback failure ("pause"|"continue") (default "pause")
	*/
	RollbackFailureAction *string

	/*
		Failure rate to tolerate during a rollback (default 0)
	*/
	RollbackMaxFailureRatio *string

	/*
		Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h) (default 5s)
	*/
	RollbackMonitor *string

	/*
		Rollback order ("start-first"|"stop-first") (default "stop-first")
	*/
	RollbackOrder *string

	/*
		Maximum number of tasks rolled back simultaneously (0 to roll back all at once)
	*/
	RollbackParallelism *uint64

	/*
		Specify secrets to expose to the service
	*/
	Secret *string

	/*
		Time to wait before force killing a container (ns|us|ms|s|m|h) (default 10s)
	*/
	StopGracePeriod *string

	/*
		Signal to stop the container
	*/
	StopSignal *string

	/*
		Sysctl options
	*/
	Sysctl []string

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Ulimit options
	*/
	Ulimit *string

	/*
		Delay between updates (ns|us|ms|s|m|h) (default 0s)
	*/
	UpdateDelay *string

	/*
		Action on update failure ("pause"|"continue"|"rollback") (default "pause")
	*/
	UpdateFailureAction *string

	/*
		Failure rate to tolerate during an update (default 0)
	*/
	UpdateMaxFailureRatio *string

	/*
		Duration after each task update to monitor for failure (ns|us|ms|s|m|h) (default 5s)
	*/
	UpdateMonitor *string

	/*
		Update order ("start-first"|"stop-first") (default "stop-first")
	*/
	UpdateOrder *string

	/*
		Maximum number of tasks updated simultaneously (0 to update all at once)
	*/
	UpdateParallelism *uint64

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		Send registry authentication details to swarm agents
	*/
	WithRegistryAuth *bool

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerServiceCreateCmd is wrapper of 'docker service create'
------------------------------
create [OPTIONS] IMAGE [COMMAND] [ARG...]
Create a new service
------------------------------
*/
func DockerServiceCreateCmd(opt DockerServiceCreateOption, args []string) *exec.Cmd {
	cargs := []string{"service", "create"}
	if opt.CapAdd != nil {
		for _, str := range opt.CapAdd {
			cargs = append(cargs, "--cap-add")
			cargs = append(cargs, str)
		}
	}
	if opt.CapDrop != nil {
		for _, str := range opt.CapDrop {
			cargs = append(cargs, "--cap-drop")
			cargs = append(cargs, str)
		}
	}
	if opt.Config != nil {
		cargs = append(cargs, "--config="+fmt.Sprint(*opt.Config))
	}
	if opt.Constraint != nil {
		for _, str := range opt.Constraint {
			cargs = append(cargs, "--constraint")
			cargs = append(cargs, str)
		}
	}
	if opt.ContainerLabel != nil {
		for _, str := range opt.ContainerLabel {
			cargs = append(cargs, "--container-label")
			cargs = append(cargs, str)
		}
	}
	if opt.CredentialSpec != nil {
		cargs = append(cargs, "--credential-spec="+fmt.Sprint(*opt.CredentialSpec))
	}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.Dns != nil {
		for _, str := range opt.Dns {
			cargs = append(cargs, "--dns")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOption != nil {
		for _, str := range opt.DnsOption {
			cargs = append(cargs, "--dns-option")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsSearch != nil {
		for _, str := range opt.DnsSearch {
			cargs = append(cargs, "--dns-search")
			cargs = append(cargs, str)
		}
	}
	if opt.EndpointMode != nil {
		cargs = append(cargs, "--endpoint-mode="+fmt.Sprint(*opt.EndpointMode))
	}
	if opt.Entrypoint != nil {
		cargs = append(cargs, "--entrypoint="+fmt.Sprint(*opt.Entrypoint))
	}
	if opt.Env != nil {
		for _, str := range opt.Env {
			cargs = append(cargs, "--env")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvFile != nil {
		for _, str := range opt.EnvFile {
			cargs = append(cargs, "--env-file")
			cargs = append(cargs, str)
		}
	}
	if opt.GenericResource != nil {
		for _, str := range opt.GenericResource {
			cargs = append(cargs, "--generic-resource")
			cargs = append(cargs, str)
		}
	}
	if opt.Group != nil {
		for _, str := range opt.Group {
			cargs = append(cargs, "--group")
			cargs = append(cargs, str)
		}
	}
	if opt.HealthCmd != nil {
		cargs = append(cargs, "--health-cmd="+fmt.Sprint(*opt.HealthCmd))
	}
	if opt.HealthInterval != nil {
		cargs = append(cargs, "--health-interval="+fmt.Sprint(*opt.HealthInterval))
	}
	if opt.HealthRetries != nil {
		cargs = append(cargs, "--health-retries="+fmt.Sprint(*opt.HealthRetries))
	}
	if opt.HealthStartPeriod != nil {
		cargs = append(cargs, "--health-start-period="+fmt.Sprint(*opt.HealthStartPeriod))
	}
	if opt.HealthTimeout != nil {
		cargs = append(cargs, "--health-timeout="+fmt.Sprint(*opt.HealthTimeout))
	}
	if opt.Host != nil {
		for _, str := range opt.Host {
			cargs = append(cargs, "--host")
			cargs = append(cargs, str)
		}
	}
	if opt.Hostname != nil {
		cargs = append(cargs, "--hostname="+fmt.Sprint(*opt.Hostname))
	}
	if opt.Init != nil {
		cargs = append(cargs, "--init="+fmt.Sprint(*opt.Init))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.LimitCpu != nil {
		cargs = append(cargs, "--limit-cpu="+fmt.Sprint(*opt.LimitCpu))
	}
	if opt.LimitMemory != nil {
		cargs = append(cargs, "--limit-memory="+fmt.Sprint(*opt.LimitMemory))
	}
	if opt.LimitPids != nil {
		cargs = append(cargs, "--limit-pids="+fmt.Sprint(*opt.LimitPids))
	}
	if opt.LogDriver != nil {
		cargs = append(cargs, "--log-driver="+fmt.Sprint(*opt.LogDriver))
	}
	if opt.LogOpt != nil {
		for _, str := range opt.LogOpt {
			cargs = append(cargs, "--log-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.MaxConcurrent != nil {
		cargs = append(cargs, "--max-concurrent="+fmt.Sprint(*opt.MaxConcurrent))
	}
	if opt.Mode != nil {
		cargs = append(cargs, "--mode="+fmt.Sprint(*opt.Mode))
	}
	if opt.Mount != nil {
		cargs = append(cargs, "--mount="+fmt.Sprint(*opt.Mount))
	}
	if opt.Name != nil {
		cargs = append(cargs, "--name="+fmt.Sprint(*opt.Name))
	}
	if opt.Network != nil {
		cargs = append(cargs, "--network="+fmt.Sprint(*opt.Network))
	}
	if opt.NoHealthcheck != nil {
		cargs = append(cargs, "--no-healthcheck="+fmt.Sprint(*opt.NoHealthcheck))
	}
	if opt.NoResolveImage != nil {
		cargs = append(cargs, "--no-resolve-image="+fmt.Sprint(*opt.NoResolveImage))
	}
	if opt.PlacementPref != nil {
		cargs = append(cargs, "--placement-pref="+fmt.Sprint(*opt.PlacementPref))
	}
	if opt.Publish != nil {
		cargs = append(cargs, "--publish="+fmt.Sprint(*opt.Publish))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.ReadOnly != nil {
		cargs = append(cargs, "--read-only="+fmt.Sprint(*opt.ReadOnly))
	}
	if opt.Replicas != nil {
		cargs = append(cargs, "--replicas="+fmt.Sprint(*opt.Replicas))
	}
	if opt.ReplicasMaxPerNode != nil {
		cargs = append(cargs, "--replicas-max-per-node="+fmt.Sprint(*opt.ReplicasMaxPerNode))
	}
	if opt.ReserveCpu != nil {
		cargs = append(cargs, "--reserve-cpu="+fmt.Sprint(*opt.ReserveCpu))
	}
	if opt.ReserveMemory != nil {
		cargs = append(cargs, "--reserve-memory="+fmt.Sprint(*opt.ReserveMemory))
	}
	if opt.RestartCondition != nil {
		cargs = append(cargs, "--restart-condition="+fmt.Sprint(*opt.RestartCondition))
	}
	if opt.RestartDelay != nil {
		cargs = append(cargs, "--restart-delay="+fmt.Sprint(*opt.RestartDelay))
	}
	if opt.RestartMaxAttempts != nil {
		cargs = append(cargs, "--restart-max-attempts="+fmt.Sprint(*opt.RestartMaxAttempts))
	}
	if opt.RestartWindow != nil {
		cargs = append(cargs, "--restart-window="+fmt.Sprint(*opt.RestartWindow))
	}
	if opt.RollbackDelay != nil {
		cargs = append(cargs, "--rollback-delay="+fmt.Sprint(*opt.RollbackDelay))
	}
	if opt.RollbackFailureAction != nil {
		cargs = append(cargs, "--rollback-failure-action="+fmt.Sprint(*opt.RollbackFailureAction))
	}
	if opt.RollbackMaxFailureRatio != nil {
		cargs = append(cargs, "--rollback-max-failure-ratio="+fmt.Sprint(*opt.RollbackMaxFailureRatio))
	}
	if opt.RollbackMonitor != nil {
		cargs = append(cargs, "--rollback-monitor="+fmt.Sprint(*opt.RollbackMonitor))
	}
	if opt.RollbackOrder != nil {
		cargs = append(cargs, "--rollback-order="+fmt.Sprint(*opt.RollbackOrder))
	}
	if opt.RollbackParallelism != nil {
		cargs = append(cargs, "--rollback-parallelism="+fmt.Sprint(*opt.RollbackParallelism))
	}
	if opt.Secret != nil {
		cargs = append(cargs, "--secret="+fmt.Sprint(*opt.Secret))
	}
	if opt.StopGracePeriod != nil {
		cargs = append(cargs, "--stop-grace-period="+fmt.Sprint(*opt.StopGracePeriod))
	}
	if opt.StopSignal != nil {
		cargs = append(cargs, "--stop-signal="+fmt.Sprint(*opt.StopSignal))
	}
	if opt.Sysctl != nil {
		for _, str := range opt.Sysctl {
			cargs = append(cargs, "--sysctl")
			cargs = append(cargs, str)
		}
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.Ulimit != nil {
		cargs = append(cargs, "--ulimit="+fmt.Sprint(*opt.Ulimit))
	}
	if opt.UpdateDelay != nil {
		cargs = append(cargs, "--update-delay="+fmt.Sprint(*opt.UpdateDelay))
	}
	if opt.UpdateFailureAction != nil {
		cargs = append(cargs, "--update-failure-action="+fmt.Sprint(*opt.UpdateFailureAction))
	}
	if opt.UpdateMaxFailureRatio != nil {
		cargs = append(cargs, "--update-max-failure-ratio="+fmt.Sprint(*opt.UpdateMaxFailureRatio))
	}
	if opt.UpdateMonitor != nil {
		cargs = append(cargs, "--update-monitor="+fmt.Sprint(*opt.UpdateMonitor))
	}
	if opt.UpdateOrder != nil {
		cargs = append(cargs, "--update-order="+fmt.Sprint(*opt.UpdateOrder))
	}
	if opt.UpdateParallelism != nil {
		cargs = append(cargs, "--update-parallelism="+fmt.Sprint(*opt.UpdateParallelism))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.WithRegistryAuth != nil {
		cargs = append(cargs, "--with-registry-auth="+fmt.Sprint(*opt.WithRegistryAuth))
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServiceInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Print the information in a human friendly format
	*/
	Pretty *bool
}

/*
DockerServiceInspectCmd is wrapper of 'docker service inspect'
------------------------------
inspect [OPTIONS] SERVICE [SERVICE...]
Display detailed information on one or more services
------------------------------
*/
func DockerServiceInspectCmd(opt DockerServiceInspectOption, args []string) *exec.Cmd {
	cargs := []string{"service", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Pretty != nil {
		cargs = append(cargs, "--pretty="+fmt.Sprint(*opt.Pretty))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServiceLogsOption struct {
	/*
		Show extra details provided to logs
	*/
	Details *bool

	/*
		Follow log output
	*/
	Follow *bool

	/*
		Do not map IDs to Names in output
	*/
	NoResolve *bool

	/*
		Do not include task IDs in output
	*/
	NoTaskIds *bool

	/*
		Do not truncate output
	*/
	NoTrunc *bool

	/*
		Do not neatly format logs
	*/
	Raw *bool

	/*
		Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
	*/
	Since *string

	/*
		Number of lines to show from the end of the logs
	*/
	Tail *string

	/*
		Show timestamps
	*/
	Timestamps *bool
}

/*
DockerServiceLogsCmd is wrapper of 'docker service logs'
------------------------------
logs [OPTIONS] SERVICE|TASK
Fetch the logs of a service or task
------------------------------
*/
func DockerServiceLogsCmd(opt DockerServiceLogsOption, args []string) *exec.Cmd {
	cargs := []string{"service", "logs"}
	if opt.Details != nil {
		cargs = append(cargs, "--details="+fmt.Sprint(*opt.Details))
	}
	if opt.Follow != nil {
		cargs = append(cargs, "--follow="+fmt.Sprint(*opt.Follow))
	}
	if opt.NoResolve != nil {
		cargs = append(cargs, "--no-resolve="+fmt.Sprint(*opt.NoResolve))
	}
	if opt.NoTaskIds != nil {
		cargs = append(cargs, "--no-task-ids="+fmt.Sprint(*opt.NoTaskIds))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Raw != nil {
		cargs = append(cargs, "--raw="+fmt.Sprint(*opt.Raw))
	}
	if opt.Since != nil {
		cargs = append(cargs, "--since="+fmt.Sprint(*opt.Since))
	}
	if opt.Tail != nil {
		cargs = append(cargs, "--tail="+fmt.Sprint(*opt.Tail))
	}
	if opt.Timestamps != nil {
		cargs = append(cargs, "--timestamps="+fmt.Sprint(*opt.Timestamps))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServiceLsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print services using a Go template
	*/
	Format *string

	/*
		Only display IDs
	*/
	Quiet *bool
}

/*
DockerServiceLsCmd is wrapper of 'docker service ls'
------------------------------
ls [OPTIONS]
List services
------------------------------
*/
func DockerServiceLsCmd(opt DockerServiceLsOption, args []string) *exec.Cmd {
	cargs := []string{"service", "ls"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServicePsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print tasks using a Go template
	*/
	Format *string

	/*
		Do not map IDs to Names
	*/
	NoResolve *bool

	/*
		Do not truncate output
	*/
	NoTrunc *bool

	/*
		Only display task IDs
	*/
	Quiet *bool
}

/*
DockerServicePsCmd is wrapper of 'docker service ps'
------------------------------
ps [OPTIONS] SERVICE [SERVICE...]
List the tasks of one or more services
------------------------------
*/
func DockerServicePsCmd(opt DockerServicePsOption, args []string) *exec.Cmd {
	cargs := []string{"service", "ps"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoResolve != nil {
		cargs = append(cargs, "--no-resolve="+fmt.Sprint(*opt.NoResolve))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerServiceRmCmd is wrapper of 'docker service rm'
------------------------------
rm SERVICE [SERVICE...]
Remove one or more services
------------------------------
*/
func DockerServiceRmCmd(args []string) *exec.Cmd {
	cargs := []string{"service", "rm"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServiceRollbackOption struct {
	/*
		Exit immediately instead of waiting for the service to converge
	*/
	Detach *bool

	/*
		Suppress progress output
	*/
	Quiet *bool
}

/*
DockerServiceRollbackCmd is wrapper of 'docker service rollback'
------------------------------
rollback [OPTIONS] SERVICE
Revert changes to a service's configuration
------------------------------
*/
func DockerServiceRollbackCmd(opt DockerServiceRollbackOption, args []string) *exec.Cmd {
	cargs := []string{"service", "rollback"}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServiceScaleOption struct {
	/*
		Exit immediately instead of waiting for the service to converge
	*/
	Detach *bool
}

/*
DockerServiceScaleCmd is wrapper of 'docker service scale'
------------------------------
scale SERVICE=REPLICAS [SERVICE=REPLICAS...]
Scale one or multiple replicated services
------------------------------
*/
func DockerServiceScaleCmd(opt DockerServiceScaleOption, args []string) *exec.Cmd {
	cargs := []string{"service", "scale"}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerServiceUpdateOption struct {
	/*
		Service command args
	*/
	Args *string

	/*
		Add Linux capabilities
	*/
	CapAdd []string

	/*
		Drop Linux capabilities
	*/
	CapDrop []string

	/*
		Add or update a config file on a service
	*/
	ConfigAdd *string

	/*
		Remove a configuration file
	*/
	ConfigRm []string

	/*
		Add or update a placement constraint
	*/
	ConstraintAdd []string

	/*
		Remove a constraint
	*/
	ConstraintRm []string

	/*
		Add or update a container label
	*/
	ContainerLabelAdd []string

	/*
		Remove a container label by its key
	*/
	ContainerLabelRm []string

	/*
		Credential spec for managed service account (Windows only)
	*/
	CredentialSpec *string

	/*
		Exit immediately instead of waiting for the service to converge
	*/
	Detach *bool

	/*
		Add or update a custom DNS server
	*/
	DnsAdd []string

	/*
		Add or update a DNS option
	*/
	DnsOptionAdd []string

	/*
		Remove a DNS option
	*/
	DnsOptionRm []string

	/*
		Remove a custom DNS server
	*/
	DnsRm []string

	/*
		Add or update a custom DNS search domain
	*/
	DnsSearchAdd []string

	/*
		Remove a DNS search domain
	*/
	DnsSearchRm []string

	/*
		Endpoint mode (vip or dnsrr)
	*/
	EndpointMode *string

	/*
		Overwrite the default ENTRYPOINT of the image
	*/
	Entrypoint *string

	/*
		Add or update an environment variable
	*/
	EnvAdd []string

	/*
		Remove an environment variable
	*/
	EnvRm []string

	/*
		Force update even if no changes require it
	*/
	Force *bool

	/*
		Add a Generic resource
	*/
	GenericResourceAdd []string

	/*
		Remove a Generic resource
	*/
	GenericResourceRm []string

	/*
		Add an additional supplementary user group to the container
	*/
	GroupAdd []string

	/*
		Remove a previously added supplementary user group from the container
	*/
	GroupRm []string

	/*
		Command to run to check health
	*/
	HealthCmd *string

	/*
		Time between running the check (ms|s|m|h)
	*/
	HealthInterval *string

	/*
		Consecutive failures needed to report unhealthy
	*/
	HealthRetries *int

	/*
		Start period for the container to initialize before counting retries towards unstable (ms|s|m|h)
	*/
	HealthStartPeriod *string

	/*
		Maximum time to allow one check to run (ms|s|m|h)
	*/
	HealthTimeout *string

	/*
		Add a custom host-to-IP mapping (host:ip)
	*/
	HostAdd []string

	/*
		Remove a custom host-to-IP mapping (host:ip)
	*/
	HostRm []string

	/*
		Container hostname
	*/
	Hostname *string

	/*
		Service image tag
	*/
	Image *string

	/*
		Use an init inside each service container to forward signals and reap processes
	*/
	Init *bool

	/*
		Service container isolation mode
	*/
	Isolation *string

	/*
		Add or update a service label
	*/
	LabelAdd []string

	/*
		Remove a label by its key
	*/
	LabelRm []string

	/*
		Limit CPUs
	*/
	LimitCpu *string

	/*
		Limit Memory
	*/
	LimitMemory *string

	/*
		Limit maximum number of processes (default 0 = unlimited)
	*/
	LimitPids *int64

	/*
		Logging driver for service
	*/
	LogDriver *string

	/*
		Logging driver options
	*/
	LogOpt []string

	/*
		Number of job tasks to run concurrently (default equal to --replicas)
	*/
	MaxConcurrent *uint

	/*
		Add or update a mount on a service
	*/
	MountAdd *string

	/*
		Remove a mount by its target path
	*/
	MountRm []string

	/*
		Add a network
	*/
	NetworkAdd *string

	/*
		Remove a network
	*/
	NetworkRm []string

	/*
		Disable any container-specified HEALTHCHECK
	*/
	NoHealthcheck *bool

	/*
		Do not query the registry to resolve image digest and supported platforms
	*/
	NoResolveImage *bool

	/*
		Add a placement preference
	*/
	PlacementPrefAdd *string

	/*
		Remove a placement preference
	*/
	PlacementPrefRm *string

	/*
		Add or update a published port
	*/
	PublishAdd *string

	/*
		Remove a published port by its target port
	*/
	PublishRm *string

	/*
		Suppress progress output
	*/
	Quiet *bool

	/*
		Mount the container's root filesystem as read only
	*/
	ReadOnly *bool

	/*
		Number of tasks
	*/
	Replicas *uint

	/*
		Maximum number of tasks per node (default 0 = unlimited)
	*/
	ReplicasMaxPerNode *uint64

	/*
		Reserve CPUs
	*/
	ReserveCpu *string

	/*
		Reserve Memory
	*/
	ReserveMemory *string

	/*
		Restart when condition is met ("none"|"on-failure"|"any")
	*/
	RestartCondition *string

	/*
		Delay between restart attempts (ns|us|ms|s|m|h)
	*/
	RestartDelay *string

	/*
		Maximum number of restarts before giving up
	*/
	RestartMaxAttempts *uint

	/*
		Window used to evaluate the restart policy (ns|us|ms|s|m|h)
	*/
	RestartWindow *string

	/*
		Rollback to previous specification
	*/
	Rollback *bool

	/*
		Delay between task rollbacks (ns|us|ms|s|m|h)
	*/
	RollbackDelay *string

	/*
		Action on rollback failure ("pause"|"continue")
	*/
	RollbackFailureAction *string

	/*
		Failure rate to tolerate during a rollback
	*/
	RollbackMaxFailureRatio *string

	/*
		Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h)
	*/
	RollbackMonitor *string

	/*
		Rollback order ("start-first"|"stop-first")
	*/
	RollbackOrder *string

	/*
		Maximum number of tasks rolled back simultaneously (0 to roll back all at once)
	*/
	RollbackParallelism *uint64

	/*
		Add or update a secret on a service
	*/
	SecretAdd *string

	/*
		Remove a secret
	*/
	SecretRm []string

	/*
		Time to wait before force killing a container (ns|us|ms|s|m|h)
	*/
	StopGracePeriod *string

	/*
		Signal to stop the container
	*/
	StopSignal *string

	/*
		Add or update a Sysctl option
	*/
	SysctlAdd []string

	/*
		Remove a Sysctl option
	*/
	SysctlRm []string

	/*
		Allocate a pseudo-TTY
	*/
	Tty *bool

	/*
		Add or update a ulimit option
	*/
	UlimitAdd *string

	/*
		Remove a ulimit option
	*/
	UlimitRm []string

	/*
		Delay between updates (ns|us|ms|s|m|h)
	*/
	UpdateDelay *string

	/*
		Action on update failure ("pause"|"continue"|"rollback")
	*/
	UpdateFailureAction *string

	/*
		Failure rate to tolerate during an update
	*/
	UpdateMaxFailureRatio *string

	/*
		Duration after each task update to monitor for failure (ns|us|ms|s|m|h)
	*/
	UpdateMonitor *string

	/*
		Update order ("start-first"|"stop-first")
	*/
	UpdateOrder *string

	/*
		Maximum number of tasks updated simultaneously (0 to update all at once)
	*/
	UpdateParallelism *uint64

	/*
		Username or UID (format: <name|uid>[:<group|gid>])
	*/
	User *string

	/*
		Send registry authentication details to swarm agents
	*/
	WithRegistryAuth *bool

	/*
		Working directory inside the container
	*/
	Workdir *string
}

/*
DockerServiceUpdateCmd is wrapper of 'docker service update'
------------------------------
update [OPTIONS] SERVICE
Update a service
------------------------------
*/
func DockerServiceUpdateCmd(opt DockerServiceUpdateOption, args []string) *exec.Cmd {
	cargs := []string{"service", "update"}
	if opt.Args != nil {
		cargs = append(cargs, "--args="+fmt.Sprint(*opt.Args))
	}
	if opt.CapAdd != nil {
		for _, str := range opt.CapAdd {
			cargs = append(cargs, "--cap-add")
			cargs = append(cargs, str)
		}
	}
	if opt.CapDrop != nil {
		for _, str := range opt.CapDrop {
			cargs = append(cargs, "--cap-drop")
			cargs = append(cargs, str)
		}
	}
	if opt.ConfigAdd != nil {
		cargs = append(cargs, "--config-add="+fmt.Sprint(*opt.ConfigAdd))
	}
	if opt.ConfigRm != nil {
		for _, str := range opt.ConfigRm {
			cargs = append(cargs, "--config-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.ConstraintAdd != nil {
		for _, str := range opt.ConstraintAdd {
			cargs = append(cargs, "--constraint-add")
			cargs = append(cargs, str)
		}
	}
	if opt.ConstraintRm != nil {
		for _, str := range opt.ConstraintRm {
			cargs = append(cargs, "--constraint-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.ContainerLabelAdd != nil {
		for _, str := range opt.ContainerLabelAdd {
			cargs = append(cargs, "--container-label-add")
			cargs = append(cargs, str)
		}
	}
	if opt.ContainerLabelRm != nil {
		for _, str := range opt.ContainerLabelRm {
			cargs = append(cargs, "--container-label-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.CredentialSpec != nil {
		cargs = append(cargs, "--credential-spec="+fmt.Sprint(*opt.CredentialSpec))
	}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.DnsAdd != nil {
		for _, str := range opt.DnsAdd {
			cargs = append(cargs, "--dns-add")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOptionAdd != nil {
		for _, str := range opt.DnsOptionAdd {
			cargs = append(cargs, "--dns-option-add")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsOptionRm != nil {
		for _, str := range opt.DnsOptionRm {
			cargs = append(cargs, "--dns-option-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsRm != nil {
		for _, str := range opt.DnsRm {
			cargs = append(cargs, "--dns-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsSearchAdd != nil {
		for _, str := range opt.DnsSearchAdd {
			cargs = append(cargs, "--dns-search-add")
			cargs = append(cargs, str)
		}
	}
	if opt.DnsSearchRm != nil {
		for _, str := range opt.DnsSearchRm {
			cargs = append(cargs, "--dns-search-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.EndpointMode != nil {
		cargs = append(cargs, "--endpoint-mode="+fmt.Sprint(*opt.EndpointMode))
	}
	if opt.Entrypoint != nil {
		cargs = append(cargs, "--entrypoint="+fmt.Sprint(*opt.Entrypoint))
	}
	if opt.EnvAdd != nil {
		for _, str := range opt.EnvAdd {
			cargs = append(cargs, "--env-add")
			cargs = append(cargs, str)
		}
	}
	if opt.EnvRm != nil {
		for _, str := range opt.EnvRm {
			cargs = append(cargs, "--env-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	if opt.GenericResourceAdd != nil {
		for _, str := range opt.GenericResourceAdd {
			cargs = append(cargs, "--generic-resource-add")
			cargs = append(cargs, str)
		}
	}
	if opt.GenericResourceRm != nil {
		for _, str := range opt.GenericResourceRm {
			cargs = append(cargs, "--generic-resource-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.GroupAdd != nil {
		for _, str := range opt.GroupAdd {
			cargs = append(cargs, "--group-add")
			cargs = append(cargs, str)
		}
	}
	if opt.GroupRm != nil {
		for _, str := range opt.GroupRm {
			cargs = append(cargs, "--group-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.HealthCmd != nil {
		cargs = append(cargs, "--health-cmd="+fmt.Sprint(*opt.HealthCmd))
	}
	if opt.HealthInterval != nil {
		cargs = append(cargs, "--health-interval="+fmt.Sprint(*opt.HealthInterval))
	}
	if opt.HealthRetries != nil {
		cargs = append(cargs, "--health-retries="+fmt.Sprint(*opt.HealthRetries))
	}
	if opt.HealthStartPeriod != nil {
		cargs = append(cargs, "--health-start-period="+fmt.Sprint(*opt.HealthStartPeriod))
	}
	if opt.HealthTimeout != nil {
		cargs = append(cargs, "--health-timeout="+fmt.Sprint(*opt.HealthTimeout))
	}
	if opt.HostAdd != nil {
		for _, str := range opt.HostAdd {
			cargs = append(cargs, "--host-add")
			cargs = append(cargs, str)
		}
	}
	if opt.HostRm != nil {
		for _, str := range opt.HostRm {
			cargs = append(cargs, "--host-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.Hostname != nil {
		cargs = append(cargs, "--hostname="+fmt.Sprint(*opt.Hostname))
	}
	if opt.Image != nil {
		cargs = append(cargs, "--image="+fmt.Sprint(*opt.Image))
	}
	if opt.Init != nil {
		cargs = append(cargs, "--init="+fmt.Sprint(*opt.Init))
	}
	if opt.Isolation != nil {
		cargs = append(cargs, "--isolation="+fmt.Sprint(*opt.Isolation))
	}
	if opt.LabelAdd != nil {
		for _, str := range opt.LabelAdd {
			cargs = append(cargs, "--label-add")
			cargs = append(cargs, str)
		}
	}
	if opt.LabelRm != nil {
		for _, str := range opt.LabelRm {
			cargs = append(cargs, "--label-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.LimitCpu != nil {
		cargs = append(cargs, "--limit-cpu="+fmt.Sprint(*opt.LimitCpu))
	}
	if opt.LimitMemory != nil {
		cargs = append(cargs, "--limit-memory="+fmt.Sprint(*opt.LimitMemory))
	}
	if opt.LimitPids != nil {
		cargs = append(cargs, "--limit-pids="+fmt.Sprint(*opt.LimitPids))
	}
	if opt.LogDriver != nil {
		cargs = append(cargs, "--log-driver="+fmt.Sprint(*opt.LogDriver))
	}
	if opt.LogOpt != nil {
		for _, str := range opt.LogOpt {
			cargs = append(cargs, "--log-opt")
			cargs = append(cargs, str)
		}
	}
	if opt.MaxConcurrent != nil {
		cargs = append(cargs, "--max-concurrent="+fmt.Sprint(*opt.MaxConcurrent))
	}
	if opt.MountAdd != nil {
		cargs = append(cargs, "--mount-add="+fmt.Sprint(*opt.MountAdd))
	}
	if opt.MountRm != nil {
		for _, str := range opt.MountRm {
			cargs = append(cargs, "--mount-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.NetworkAdd != nil {
		cargs = append(cargs, "--network-add="+fmt.Sprint(*opt.NetworkAdd))
	}
	if opt.NetworkRm != nil {
		for _, str := range opt.NetworkRm {
			cargs = append(cargs, "--network-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.NoHealthcheck != nil {
		cargs = append(cargs, "--no-healthcheck="+fmt.Sprint(*opt.NoHealthcheck))
	}
	if opt.NoResolveImage != nil {
		cargs = append(cargs, "--no-resolve-image="+fmt.Sprint(*opt.NoResolveImage))
	}
	if opt.PlacementPrefAdd != nil {
		cargs = append(cargs, "--placement-pref-add="+fmt.Sprint(*opt.PlacementPrefAdd))
	}
	if opt.PlacementPrefRm != nil {
		cargs = append(cargs, "--placement-pref-rm="+fmt.Sprint(*opt.PlacementPrefRm))
	}
	if opt.PublishAdd != nil {
		cargs = append(cargs, "--publish-add="+fmt.Sprint(*opt.PublishAdd))
	}
	if opt.PublishRm != nil {
		cargs = append(cargs, "--publish-rm="+fmt.Sprint(*opt.PublishRm))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.ReadOnly != nil {
		cargs = append(cargs, "--read-only="+fmt.Sprint(*opt.ReadOnly))
	}
	if opt.Replicas != nil {
		cargs = append(cargs, "--replicas="+fmt.Sprint(*opt.Replicas))
	}
	if opt.ReplicasMaxPerNode != nil {
		cargs = append(cargs, "--replicas-max-per-node="+fmt.Sprint(*opt.ReplicasMaxPerNode))
	}
	if opt.ReserveCpu != nil {
		cargs = append(cargs, "--reserve-cpu="+fmt.Sprint(*opt.ReserveCpu))
	}
	if opt.ReserveMemory != nil {
		cargs = append(cargs, "--reserve-memory="+fmt.Sprint(*opt.ReserveMemory))
	}
	if opt.RestartCondition != nil {
		cargs = append(cargs, "--restart-condition="+fmt.Sprint(*opt.RestartCondition))
	}
	if opt.RestartDelay != nil {
		cargs = append(cargs, "--restart-delay="+fmt.Sprint(*opt.RestartDelay))
	}
	if opt.RestartMaxAttempts != nil {
		cargs = append(cargs, "--restart-max-attempts="+fmt.Sprint(*opt.RestartMaxAttempts))
	}
	if opt.RestartWindow != nil {
		cargs = append(cargs, "--restart-window="+fmt.Sprint(*opt.RestartWindow))
	}
	if opt.Rollback != nil {
		cargs = append(cargs, "--rollback="+fmt.Sprint(*opt.Rollback))
	}
	if opt.RollbackDelay != nil {
		cargs = append(cargs, "--rollback-delay="+fmt.Sprint(*opt.RollbackDelay))
	}
	if opt.RollbackFailureAction != nil {
		cargs = append(cargs, "--rollback-failure-action="+fmt.Sprint(*opt.RollbackFailureAction))
	}
	if opt.RollbackMaxFailureRatio != nil {
		cargs = append(cargs, "--rollback-max-failure-ratio="+fmt.Sprint(*opt.RollbackMaxFailureRatio))
	}
	if opt.RollbackMonitor != nil {
		cargs = append(cargs, "--rollback-monitor="+fmt.Sprint(*opt.RollbackMonitor))
	}
	if opt.RollbackOrder != nil {
		cargs = append(cargs, "--rollback-order="+fmt.Sprint(*opt.RollbackOrder))
	}
	if opt.RollbackParallelism != nil {
		cargs = append(cargs, "--rollback-parallelism="+fmt.Sprint(*opt.RollbackParallelism))
	}
	if opt.SecretAdd != nil {
		cargs = append(cargs, "--secret-add="+fmt.Sprint(*opt.SecretAdd))
	}
	if opt.SecretRm != nil {
		for _, str := range opt.SecretRm {
			cargs = append(cargs, "--secret-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.StopGracePeriod != nil {
		cargs = append(cargs, "--stop-grace-period="+fmt.Sprint(*opt.StopGracePeriod))
	}
	if opt.StopSignal != nil {
		cargs = append(cargs, "--stop-signal="+fmt.Sprint(*opt.StopSignal))
	}
	if opt.SysctlAdd != nil {
		for _, str := range opt.SysctlAdd {
			cargs = append(cargs, "--sysctl-add")
			cargs = append(cargs, str)
		}
	}
	if opt.SysctlRm != nil {
		for _, str := range opt.SysctlRm {
			cargs = append(cargs, "--sysctl-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.Tty != nil {
		cargs = append(cargs, "--tty="+fmt.Sprint(*opt.Tty))
	}
	if opt.UlimitAdd != nil {
		cargs = append(cargs, "--ulimit-add="+fmt.Sprint(*opt.UlimitAdd))
	}
	if opt.UlimitRm != nil {
		for _, str := range opt.UlimitRm {
			cargs = append(cargs, "--ulimit-rm")
			cargs = append(cargs, str)
		}
	}
	if opt.UpdateDelay != nil {
		cargs = append(cargs, "--update-delay="+fmt.Sprint(*opt.UpdateDelay))
	}
	if opt.UpdateFailureAction != nil {
		cargs = append(cargs, "--update-failure-action="+fmt.Sprint(*opt.UpdateFailureAction))
	}
	if opt.UpdateMaxFailureRatio != nil {
		cargs = append(cargs, "--update-max-failure-ratio="+fmt.Sprint(*opt.UpdateMaxFailureRatio))
	}
	if opt.UpdateMonitor != nil {
		cargs = append(cargs, "--update-monitor="+fmt.Sprint(*opt.UpdateMonitor))
	}
	if opt.UpdateOrder != nil {
		cargs = append(cargs, "--update-order="+fmt.Sprint(*opt.UpdateOrder))
	}
	if opt.UpdateParallelism != nil {
		cargs = append(cargs, "--update-parallelism="+fmt.Sprint(*opt.UpdateParallelism))
	}
	if opt.User != nil {
		cargs = append(cargs, "--user="+fmt.Sprint(*opt.User))
	}
	if opt.WithRegistryAuth != nil {
		cargs = append(cargs, "--with-registry-auth="+fmt.Sprint(*opt.WithRegistryAuth))
	}
	if opt.Workdir != nil {
		cargs = append(cargs, "--workdir="+fmt.Sprint(*opt.Workdir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerStackCmd is wrapper of 'docker stack'
------------------------------
stack [OPTIONS]
Manage Docker stacks
------------------------------
*/
func DockerStackCmd(args []string) *exec.Cmd {
	cargs := []string{"stack"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStackDeployOption struct {
	/*
		Path to a Compose file, or "-" to read from stdin
	*/
	ComposeFile *string

	/*
		Kubernetes namespace to use
	*/
	Namespace *string

	/*
		Prune services that are no longer referenced
	*/
	Prune *bool

	/*
		Query the registry to resolve image digest and supported platforms ("always"|"changed"|"never")
	*/
	ResolveImage *string

	/*
		Send registry authentication details to Swarm agents
	*/
	WithRegistryAuth *bool
}

/*
DockerStackDeployCmd is wrapper of 'docker stack deploy'
------------------------------
deploy [OPTIONS] STACK
Deploy a new stack or update an existing stack
------------------------------
*/
func DockerStackDeployCmd(opt DockerStackDeployOption, args []string) *exec.Cmd {
	cargs := []string{"stack", "deploy"}
	if opt.ComposeFile != nil {
		cargs = append(cargs, "--compose-file="+fmt.Sprint(*opt.ComposeFile))
	}
	if opt.Namespace != nil {
		cargs = append(cargs, "--namespace="+fmt.Sprint(*opt.Namespace))
	}
	if opt.Prune != nil {
		cargs = append(cargs, "--prune="+fmt.Sprint(*opt.Prune))
	}
	if opt.ResolveImage != nil {
		cargs = append(cargs, "--resolve-image="+fmt.Sprint(*opt.ResolveImage))
	}
	if opt.WithRegistryAuth != nil {
		cargs = append(cargs, "--with-registry-auth="+fmt.Sprint(*opt.WithRegistryAuth))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStackLsOption struct {
	/*
		List stacks from all Kubernetes namespaces
	*/
	AllNamespaces *bool

	/*
		Pretty-print stacks using a Go template
	*/
	Format *string

	/*
		Kubernetes namespaces to use
	*/
	Namespace *string
}

/*
DockerStackLsCmd is wrapper of 'docker stack ls'
------------------------------
ls [OPTIONS]
List stacks
------------------------------
*/
func DockerStackLsCmd(opt DockerStackLsOption, args []string) *exec.Cmd {
	cargs := []string{"stack", "ls"}
	if opt.AllNamespaces != nil {
		cargs = append(cargs, "--all-namespaces="+fmt.Sprint(*opt.AllNamespaces))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Namespace != nil {
		cargs = append(cargs, "--namespace="+fmt.Sprint(*opt.Namespace))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStackPsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print tasks using a Go template
	*/
	Format *string

	/*
		Kubernetes namespace to use
	*/
	Namespace *string

	/*
		Do not map IDs to Names
	*/
	NoResolve *bool

	/*
		Do not truncate output
	*/
	NoTrunc *bool

	/*
		Only display task IDs
	*/
	Quiet *bool
}

/*
DockerStackPsCmd is wrapper of 'docker stack ps'
------------------------------
ps [OPTIONS] STACK
List the tasks in the stack
------------------------------
*/
func DockerStackPsCmd(opt DockerStackPsOption, args []string) *exec.Cmd {
	cargs := []string{"stack", "ps"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Namespace != nil {
		cargs = append(cargs, "--namespace="+fmt.Sprint(*opt.Namespace))
	}
	if opt.NoResolve != nil {
		cargs = append(cargs, "--no-resolve="+fmt.Sprint(*opt.NoResolve))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStackRmOption struct {
	/*
		Kubernetes namespace to use
	*/
	Namespace *string
}

/*
DockerStackRmCmd is wrapper of 'docker stack rm'
------------------------------
rm [OPTIONS] STACK [STACK...]
Remove one or more stacks
------------------------------
*/
func DockerStackRmCmd(opt DockerStackRmOption, args []string) *exec.Cmd {
	cargs := []string{"stack", "rm"}
	if opt.Namespace != nil {
		cargs = append(cargs, "--namespace="+fmt.Sprint(*opt.Namespace))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStackServicesOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Pretty-print services using a Go template
	*/
	Format *string

	/*
		Kubernetes namespace to use
	*/
	Namespace *string

	/*
		Only display IDs
	*/
	Quiet *bool
}

/*
DockerStackServicesCmd is wrapper of 'docker stack services'
------------------------------
services [OPTIONS] STACK
List the services in the stack
------------------------------
*/
func DockerStackServicesCmd(opt DockerStackServicesOption, args []string) *exec.Cmd {
	cargs := []string{"stack", "services"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Namespace != nil {
		cargs = append(cargs, "--namespace="+fmt.Sprint(*opt.Namespace))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStartOption struct {
	/*
		Attach STDOUT/STDERR and forward signals
	*/
	Attach *bool

	/*
		Restore from this checkpoint
	*/
	Checkpoint *string

	/*
		Use a custom checkpoint storage directory
	*/
	CheckpointDir *string

	/*
		Override the key sequence for detaching a container
	*/
	DetachKeys *string

	/*
		Attach container's STDIN
	*/
	Interactive *bool
}

/*
DockerStartCmd is wrapper of 'docker start'
------------------------------
start [OPTIONS] CONTAINER [CONTAINER...]
Start one or more stopped containers
------------------------------
*/
func DockerStartCmd(opt DockerStartOption, args []string) *exec.Cmd {
	cargs := []string{"start"}
	if opt.Attach != nil {
		cargs = append(cargs, "--attach="+fmt.Sprint(*opt.Attach))
	}
	if opt.Checkpoint != nil {
		cargs = append(cargs, "--checkpoint="+fmt.Sprint(*opt.Checkpoint))
	}
	if opt.CheckpointDir != nil {
		cargs = append(cargs, "--checkpoint-dir="+fmt.Sprint(*opt.CheckpointDir))
	}
	if opt.DetachKeys != nil {
		cargs = append(cargs, "--detach-keys="+fmt.Sprint(*opt.DetachKeys))
	}
	if opt.Interactive != nil {
		cargs = append(cargs, "--interactive="+fmt.Sprint(*opt.Interactive))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStatsOption struct {
	/*
		Show all containers (default shows just running)
	*/
	All *bool

	/*
		Pretty-print images using a Go template
	*/
	Format *string

	/*
		Disable streaming stats and only pull the first result
	*/
	NoStream *bool

	/*
		Do not truncate output
	*/
	NoTrunc *bool
}

/*
DockerStatsCmd is wrapper of 'docker stats'
------------------------------
stats [OPTIONS] [CONTAINER...]
Display a live stream of container(s) resource usage statistics
------------------------------
*/
func DockerStatsCmd(opt DockerStatsOption, args []string) *exec.Cmd {
	cargs := []string{"stats"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.NoStream != nil {
		cargs = append(cargs, "--no-stream="+fmt.Sprint(*opt.NoStream))
	}
	if opt.NoTrunc != nil {
		cargs = append(cargs, "--no-trunc="+fmt.Sprint(*opt.NoTrunc))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerStopOption struct {
	/*
		Seconds to wait for stop before killing it
	*/
	Time *int
}

/*
DockerStopCmd is wrapper of 'docker stop'
------------------------------
stop [OPTIONS] CONTAINER [CONTAINER...]
Stop one or more running containers
------------------------------
*/
func DockerStopCmd(opt DockerStopOption, args []string) *exec.Cmd {
	cargs := []string{"stop"}
	if opt.Time != nil {
		cargs = append(cargs, "--time="+fmt.Sprint(*opt.Time))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerSwarmCmd is wrapper of 'docker swarm'
------------------------------
swarm
Manage Swarm
------------------------------
*/
func DockerSwarmCmd(args []string) *exec.Cmd {
	cargs := []string{"swarm"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSwarmCaOption struct {
	/*
		Path to the PEM-formatted root CA certificate to use for the new cluster
	*/
	CaCert *string

	/*
		Path to the PEM-formatted root CA key to use for the new cluster
	*/
	CaKey *string

	/*
		Validity period for node certificates (ns|us|ms|s|m|h)
	*/
	CertExpiry *string

	/*
		Exit immediately instead of waiting for the root rotation to converge
	*/
	Detach *bool

	/*
		Specifications of one or more certificate signing endpoints
	*/
	ExternalCa *string

	/*
		Suppress progress output
	*/
	Quiet *bool

	/*
		Rotate the swarm CA - if no certificate or key are provided, new ones will be generated
	*/
	Rotate *bool
}

/*
DockerSwarmCaCmd is wrapper of 'docker swarm ca'
------------------------------
ca [OPTIONS]
Display and rotate the root CA
------------------------------
*/
func DockerSwarmCaCmd(opt DockerSwarmCaOption, args []string) *exec.Cmd {
	cargs := []string{"swarm", "ca"}
	if opt.CaCert != nil {
		cargs = append(cargs, "--ca-cert="+fmt.Sprint(*opt.CaCert))
	}
	if opt.CaKey != nil {
		cargs = append(cargs, "--ca-key="+fmt.Sprint(*opt.CaKey))
	}
	if opt.CertExpiry != nil {
		cargs = append(cargs, "--cert-expiry="+fmt.Sprint(*opt.CertExpiry))
	}
	if opt.Detach != nil {
		cargs = append(cargs, "--detach="+fmt.Sprint(*opt.Detach))
	}
	if opt.ExternalCa != nil {
		cargs = append(cargs, "--external-ca="+fmt.Sprint(*opt.ExternalCa))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Rotate != nil {
		cargs = append(cargs, "--rotate="+fmt.Sprint(*opt.Rotate))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSwarmInitOption struct {
	/*
		Advertised address (format: <ip|interface>[:port])
	*/
	AdvertiseAddr *string

	/*
		Enable manager autolocking (requiring an unlock key to start a stopped manager)
	*/
	Autolock *bool

	/*
		Availability of the node ("active"|"pause"|"drain")
	*/
	Availability *string

	/*
		Validity period for node certificates (ns|us|ms|s|m|h)
	*/
	CertExpiry *string

	/*
		Address or interface to use for data path traffic (format: <ip|interface>)
	*/
	DataPathAddr *string

	/*
		Port number to use for data path traffic (1024 - 49151). If no value is set or is set to 0, the default port (4789) is used.
	*/
	DataPathPort *uint32

	/*
		default address pool in CIDR format
	*/
	DefaultAddrPool *string

	/*
		default address pool subnet mask length
	*/
	DefaultAddrPoolMaskLength *uint32

	/*
		Dispatcher heartbeat period (ns|us|ms|s|m|h)
	*/
	DispatcherHeartbeat *string

	/*
		Specifications of one or more certificate signing endpoints
	*/
	ExternalCa *string

	/*
		Force create a new cluster from current state
	*/
	ForceNewCluster *bool

	/*
		Listen address (format: <ip|interface>[:port])
	*/
	ListenAddr *string

	/*
		Number of additional Raft snapshots to retain
	*/
	MaxSnapshots *uint64

	/*
		Number of log entries between Raft snapshots
	*/
	SnapshotInterval *uint64

	/*
		Task history retention limit
	*/
	TaskHistoryLimit *int64
}

/*
DockerSwarmInitCmd is wrapper of 'docker swarm init'
------------------------------
init [OPTIONS]
Initialize a swarm
------------------------------
*/
func DockerSwarmInitCmd(opt DockerSwarmInitOption, args []string) *exec.Cmd {
	cargs := []string{"swarm", "init"}
	if opt.AdvertiseAddr != nil {
		cargs = append(cargs, "--advertise-addr="+fmt.Sprint(*opt.AdvertiseAddr))
	}
	if opt.Autolock != nil {
		cargs = append(cargs, "--autolock="+fmt.Sprint(*opt.Autolock))
	}
	if opt.Availability != nil {
		cargs = append(cargs, "--availability="+fmt.Sprint(*opt.Availability))
	}
	if opt.CertExpiry != nil {
		cargs = append(cargs, "--cert-expiry="+fmt.Sprint(*opt.CertExpiry))
	}
	if opt.DataPathAddr != nil {
		cargs = append(cargs, "--data-path-addr="+fmt.Sprint(*opt.DataPathAddr))
	}
	if opt.DataPathPort != nil {
		cargs = append(cargs, "--data-path-port="+fmt.Sprint(*opt.DataPathPort))
	}
	if opt.DefaultAddrPool != nil {
		cargs = append(cargs, "--default-addr-pool="+fmt.Sprint(*opt.DefaultAddrPool))
	}
	if opt.DefaultAddrPoolMaskLength != nil {
		cargs = append(cargs, "--default-addr-pool-mask-length="+fmt.Sprint(*opt.DefaultAddrPoolMaskLength))
	}
	if opt.DispatcherHeartbeat != nil {
		cargs = append(cargs, "--dispatcher-heartbeat="+fmt.Sprint(*opt.DispatcherHeartbeat))
	}
	if opt.ExternalCa != nil {
		cargs = append(cargs, "--external-ca="+fmt.Sprint(*opt.ExternalCa))
	}
	if opt.ForceNewCluster != nil {
		cargs = append(cargs, "--force-new-cluster="+fmt.Sprint(*opt.ForceNewCluster))
	}
	if opt.ListenAddr != nil {
		cargs = append(cargs, "--listen-addr="+fmt.Sprint(*opt.ListenAddr))
	}
	if opt.MaxSnapshots != nil {
		cargs = append(cargs, "--max-snapshots="+fmt.Sprint(*opt.MaxSnapshots))
	}
	if opt.SnapshotInterval != nil {
		cargs = append(cargs, "--snapshot-interval="+fmt.Sprint(*opt.SnapshotInterval))
	}
	if opt.TaskHistoryLimit != nil {
		cargs = append(cargs, "--task-history-limit="+fmt.Sprint(*opt.TaskHistoryLimit))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSwarmJoinOption struct {
	/*
		Advertised address (format: <ip|interface>[:port])
	*/
	AdvertiseAddr *string

	/*
		Availability of the node ("active"|"pause"|"drain")
	*/
	Availability *string

	/*
		Address or interface to use for data path traffic (format: <ip|interface>)
	*/
	DataPathAddr *string

	/*
		Listen address (format: <ip|interface>[:port])
	*/
	ListenAddr *string

	/*
		Token for entry into the swarm
	*/
	Token *string
}

/*
DockerSwarmJoinCmd is wrapper of 'docker swarm join'
------------------------------
join [OPTIONS] HOST:PORT
Join a swarm as a node and/or manager
------------------------------
*/
func DockerSwarmJoinCmd(opt DockerSwarmJoinOption, args []string) *exec.Cmd {
	cargs := []string{"swarm", "join"}
	if opt.AdvertiseAddr != nil {
		cargs = append(cargs, "--advertise-addr="+fmt.Sprint(*opt.AdvertiseAddr))
	}
	if opt.Availability != nil {
		cargs = append(cargs, "--availability="+fmt.Sprint(*opt.Availability))
	}
	if opt.DataPathAddr != nil {
		cargs = append(cargs, "--data-path-addr="+fmt.Sprint(*opt.DataPathAddr))
	}
	if opt.ListenAddr != nil {
		cargs = append(cargs, "--listen-addr="+fmt.Sprint(*opt.ListenAddr))
	}
	if opt.Token != nil {
		cargs = append(cargs, "--token="+fmt.Sprint(*opt.Token))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSwarmJoinTokenOption struct {
	/*
		Only display token
	*/
	Quiet *bool

	/*
		Rotate join token
	*/
	Rotate *bool
}

/*
DockerSwarmJoinTokenCmd is wrapper of 'docker swarm join-token'
------------------------------
join-token [OPTIONS] (worker|manager)
Manage join tokens
------------------------------
*/
func DockerSwarmJoinTokenCmd(opt DockerSwarmJoinTokenOption, args []string) *exec.Cmd {
	cargs := []string{"swarm", "join-token"}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Rotate != nil {
		cargs = append(cargs, "--rotate="+fmt.Sprint(*opt.Rotate))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSwarmLeaveOption struct {
	/*
		Force this node to leave the swarm, ignoring warnings
	*/
	Force *bool
}

/*
DockerSwarmLeaveCmd is wrapper of 'docker swarm leave'
------------------------------
leave [OPTIONS]
Leave the swarm
------------------------------
*/
func DockerSwarmLeaveCmd(opt DockerSwarmLeaveOption, args []string) *exec.Cmd {
	cargs := []string{"swarm", "leave"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerSwarmUnlockCmd is wrapper of 'docker swarm unlock'
------------------------------
unlock
Unlock swarm
------------------------------
*/
func DockerSwarmUnlockCmd(args []string) *exec.Cmd {
	cargs := []string{"swarm", "unlock"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSwarmUnlockKeyOption struct {
	/*
		Only display token
	*/
	Quiet *bool

	/*
		Rotate unlock key
	*/
	Rotate *bool
}

/*
DockerSwarmUnlockKeyCmd is wrapper of 'docker swarm unlock-key'
------------------------------
unlock-key [OPTIONS]
Manage the unlock key
------------------------------
*/
func DockerSwarmUnlockKeyCmd(opt DockerSwarmUnlockKeyOption, args []string) *exec.Cmd {
	cargs := []string{"swarm", "unlock-key"}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	if opt.Rotate != nil {
		cargs = append(cargs, "--rotate="+fmt.Sprint(*opt.Rotate))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSwarmUpdateOption struct {
	/*
		Change manager autolocking setting (true|false)
	*/
	Autolock *bool

	/*
		Validity period for node certificates (ns|us|ms|s|m|h)
	*/
	CertExpiry *string

	/*
		Dispatcher heartbeat period (ns|us|ms|s|m|h)
	*/
	DispatcherHeartbeat *string

	/*
		Specifications of one or more certificate signing endpoints
	*/
	ExternalCa *string

	/*
		Number of additional Raft snapshots to retain
	*/
	MaxSnapshots *uint64

	/*
		Number of log entries between Raft snapshots
	*/
	SnapshotInterval *uint64

	/*
		Task history retention limit
	*/
	TaskHistoryLimit *int64
}

/*
DockerSwarmUpdateCmd is wrapper of 'docker swarm update'
------------------------------
update [OPTIONS]
Update the swarm
------------------------------
*/
func DockerSwarmUpdateCmd(opt DockerSwarmUpdateOption, args []string) *exec.Cmd {
	cargs := []string{"swarm", "update"}
	if opt.Autolock != nil {
		cargs = append(cargs, "--autolock="+fmt.Sprint(*opt.Autolock))
	}
	if opt.CertExpiry != nil {
		cargs = append(cargs, "--cert-expiry="+fmt.Sprint(*opt.CertExpiry))
	}
	if opt.DispatcherHeartbeat != nil {
		cargs = append(cargs, "--dispatcher-heartbeat="+fmt.Sprint(*opt.DispatcherHeartbeat))
	}
	if opt.ExternalCa != nil {
		cargs = append(cargs, "--external-ca="+fmt.Sprint(*opt.ExternalCa))
	}
	if opt.MaxSnapshots != nil {
		cargs = append(cargs, "--max-snapshots="+fmt.Sprint(*opt.MaxSnapshots))
	}
	if opt.SnapshotInterval != nil {
		cargs = append(cargs, "--snapshot-interval="+fmt.Sprint(*opt.SnapshotInterval))
	}
	if opt.TaskHistoryLimit != nil {
		cargs = append(cargs, "--task-history-limit="+fmt.Sprint(*opt.TaskHistoryLimit))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerSystemCmd is wrapper of 'docker system'
------------------------------
system
Manage Docker
------------------------------
*/
func DockerSystemCmd(args []string) *exec.Cmd {
	cargs := []string{"system"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSystemDfOption struct {
	/*
		Pretty-print images using a Go template
	*/
	Format *string

	/*
		Show detailed information on space usage
	*/
	Verbose *bool
}

/*
DockerSystemDfCmd is wrapper of 'docker system df'
------------------------------
df [OPTIONS]
Show docker disk usage
------------------------------
*/
func DockerSystemDfCmd(opt DockerSystemDfOption, args []string) *exec.Cmd {
	cargs := []string{"system", "df"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Verbose != nil {
		cargs = append(cargs, "--verbose="+fmt.Sprint(*opt.Verbose))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerSystemDialStdioCmd is wrapper of 'docker system dial-stdio'
------------------------------
dial-stdio
Proxy the stdio stream to the daemon connection. Should not be invoked manually.
------------------------------
*/
func DockerSystemDialStdioCmd(args []string) *exec.Cmd {
	cargs := []string{"system", "dial-stdio"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSystemEventsOption struct {
	/*
		Filter output based on conditions provided
	*/
	Filter *string

	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Show all events created since timestamp
	*/
	Since *string

	/*
		Stream events until this timestamp
	*/
	Until *string
}

/*
DockerSystemEventsCmd is wrapper of 'docker system events'
------------------------------
events [OPTIONS]
Get real time events from the server
------------------------------
*/
func DockerSystemEventsCmd(opt DockerSystemEventsOption, args []string) *exec.Cmd {
	cargs := []string{"system", "events"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Since != nil {
		cargs = append(cargs, "--since="+fmt.Sprint(*opt.Since))
	}
	if opt.Until != nil {
		cargs = append(cargs, "--until="+fmt.Sprint(*opt.Until))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSystemInfoOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string
}

/*
DockerSystemInfoCmd is wrapper of 'docker system info'
------------------------------
info [OPTIONS]
Display system-wide information
------------------------------
*/
func DockerSystemInfoCmd(opt DockerSystemInfoOption, args []string) *exec.Cmd {
	cargs := []string{"system", "info"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerSystemPruneOption struct {
	/*
		Remove all unused images not just dangling ones
	*/
	All *bool

	/*
		Provide filter values (e.g. 'label=<key>=<value>')
	*/
	Filter *string

	/*
		Do not prompt for confirmation
	*/
	Force *bool

	/*
		Prune volumes
	*/
	Volumes *bool
}

/*
DockerSystemPruneCmd is wrapper of 'docker system prune'
------------------------------
prune [OPTIONS]
Remove unused data
------------------------------
*/
func DockerSystemPruneCmd(opt DockerSystemPruneOption, args []string) *exec.Cmd {
	cargs := []string{"system", "prune"}
	if opt.All != nil {
		cargs = append(cargs, "--all="+fmt.Sprint(*opt.All))
	}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	if opt.Volumes != nil {
		cargs = append(cargs, "--volumes="+fmt.Sprint(*opt.Volumes))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerTagCmd is wrapper of 'docker tag'
------------------------------
tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
------------------------------
*/
func DockerTagCmd(args []string) *exec.Cmd {
	cargs := []string{"tag"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerTopCmd is wrapper of 'docker top'
------------------------------
top CONTAINER [ps OPTIONS]
Display the running processes of a container
------------------------------
*/
func DockerTopCmd(args []string) *exec.Cmd {
	cargs := []string{"top"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerTrustCmd is wrapper of 'docker trust'
------------------------------
trust
Manage trust on Docker images
------------------------------
*/
func DockerTrustCmd(args []string) *exec.Cmd {
	cargs := []string{"trust"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerTrustInspectOption struct {
	/*
		Print the information in a human friendly format
	*/
	Pretty *bool
}

/*
DockerTrustInspectCmd is wrapper of 'docker trust inspect'
------------------------------
inspect IMAGE[:TAG] [IMAGE[:TAG]...]
Return low-level information about keys and signatures
------------------------------
*/
func DockerTrustInspectCmd(opt DockerTrustInspectOption, args []string) *exec.Cmd {
	cargs := []string{"trust", "inspect"}
	if opt.Pretty != nil {
		cargs = append(cargs, "--pretty="+fmt.Sprint(*opt.Pretty))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerTrustKeyCmd is wrapper of 'docker trust key'
------------------------------
key
Manage keys for signing Docker images
------------------------------
*/
func DockerTrustKeyCmd(args []string) *exec.Cmd {
	cargs := []string{"trust", "key"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerTrustKeyGenerateOption struct {
	/*
		Directory to generate key in, defaults to current directory
	*/
	Dir *string
}

/*
DockerTrustKeyGenerateCmd is wrapper of 'docker trust key generate'
------------------------------
generate NAME
Generate and load a signing key-pair
------------------------------
*/
func DockerTrustKeyGenerateCmd(opt DockerTrustKeyGenerateOption, args []string) *exec.Cmd {
	cargs := []string{"trust", "key", "generate"}
	if opt.Dir != nil {
		cargs = append(cargs, "--dir="+fmt.Sprint(*opt.Dir))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerTrustKeyLoadOption struct {
	/*
		Name for the loaded key
	*/
	Name *string
}

/*
DockerTrustKeyLoadCmd is wrapper of 'docker trust key load'
------------------------------
load [OPTIONS] KEYFILE
Load a private key file for signing
------------------------------
*/
func DockerTrustKeyLoadCmd(opt DockerTrustKeyLoadOption, args []string) *exec.Cmd {
	cargs := []string{"trust", "key", "load"}
	if opt.Name != nil {
		cargs = append(cargs, "--name="+fmt.Sprint(*opt.Name))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerTrustRevokeOption struct {
	/*
		Do not prompt for confirmation
	*/
	Yes *bool
}

/*
DockerTrustRevokeCmd is wrapper of 'docker trust revoke'
------------------------------
revoke [OPTIONS] IMAGE[:TAG]
Remove trust for an image
------------------------------
*/
func DockerTrustRevokeCmd(opt DockerTrustRevokeOption, args []string) *exec.Cmd {
	cargs := []string{"trust", "revoke"}
	if opt.Yes != nil {
		cargs = append(cargs, "--yes="+fmt.Sprint(*opt.Yes))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerTrustSignOption struct {
	/*
		Sign a locally tagged image
	*/
	Local *bool
}

/*
DockerTrustSignCmd is wrapper of 'docker trust sign'
------------------------------
sign IMAGE:TAG
Sign an image
------------------------------
*/
func DockerTrustSignCmd(opt DockerTrustSignOption, args []string) *exec.Cmd {
	cargs := []string{"trust", "sign"}
	if opt.Local != nil {
		cargs = append(cargs, "--local="+fmt.Sprint(*opt.Local))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerTrustSignerCmd is wrapper of 'docker trust signer'
------------------------------
signer
Manage entities who can sign Docker images
------------------------------
*/
func DockerTrustSignerCmd(args []string) *exec.Cmd {
	cargs := []string{"trust", "signer"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerTrustSignerAddOption struct {
	/*
		Path to the signer's public key file
	*/
	Key []string
}

/*
DockerTrustSignerAddCmd is wrapper of 'docker trust signer add'
------------------------------
add OPTIONS NAME REPOSITORY [REPOSITORY...]
Add a signer
------------------------------
*/
func DockerTrustSignerAddCmd(opt DockerTrustSignerAddOption, args []string) *exec.Cmd {
	cargs := []string{"trust", "signer", "add"}
	if opt.Key != nil {
		for _, str := range opt.Key {
			cargs = append(cargs, "--key")
			cargs = append(cargs, str)
		}
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerTrustSignerRemoveOption struct {
	/*
		Do not prompt for confirmation before removing the most recent signer
	*/
	Force *bool
}

/*
DockerTrustSignerRemoveCmd is wrapper of 'docker trust signer remove'
------------------------------
remove [OPTIONS] NAME REPOSITORY [REPOSITORY...]
Remove a signer
------------------------------
*/
func DockerTrustSignerRemoveCmd(opt DockerTrustSignerRemoveOption, args []string) *exec.Cmd {
	cargs := []string{"trust", "signer", "remove"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerUnpauseCmd is wrapper of 'docker unpause'
------------------------------
unpause CONTAINER [CONTAINER...]
Unpause all processes within one or more containers
------------------------------
*/
func DockerUnpauseCmd(args []string) *exec.Cmd {
	cargs := []string{"unpause"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerUpdateOption struct {
	/*
		Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
	*/
	BlkioWeight *uint16

	/*
		Limit CPU CFS (Completely Fair Scheduler) period
	*/
	CpuPeriod *int64

	/*
		Limit CPU CFS (Completely Fair Scheduler) quota
	*/
	CpuQuota *int64

	/*
		Limit the CPU real-time period in microseconds
	*/
	CpuRtPeriod *int64

	/*
		Limit the CPU real-time runtime in microseconds
	*/
	CpuRtRuntime *int64

	/*
		CPU shares (relative weight)
	*/
	CpuShares *int64

	/*
		Number of CPUs
	*/
	Cpus *string

	/*
		CPUs in which to allow execution (0-3, 0,1)
	*/
	CpusetCpus *string

	/*
		MEMs in which to allow execution (0-3, 0,1)
	*/
	CpusetMems *string

	/*
		Kernel memory limit
	*/
	KernelMemory *string

	/*
		Memory limit
	*/
	Memory *string

	/*
		Memory soft limit
	*/
	MemoryReservation *string

	/*
		Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	*/
	MemorySwap *string

	/*
		Tune container pids limit (set -1 for unlimited)
	*/
	PidsLimit *int64

	/*
		Restart policy to apply when a container exits
	*/
	Restart *string
}

/*
DockerUpdateCmd is wrapper of 'docker update'
------------------------------
update [OPTIONS] CONTAINER [CONTAINER...]
Update configuration of one or more containers
------------------------------
*/
func DockerUpdateCmd(opt DockerUpdateOption, args []string) *exec.Cmd {
	cargs := []string{"update"}
	if opt.BlkioWeight != nil {
		cargs = append(cargs, "--blkio-weight="+fmt.Sprint(*opt.BlkioWeight))
	}
	if opt.CpuPeriod != nil {
		cargs = append(cargs, "--cpu-period="+fmt.Sprint(*opt.CpuPeriod))
	}
	if opt.CpuQuota != nil {
		cargs = append(cargs, "--cpu-quota="+fmt.Sprint(*opt.CpuQuota))
	}
	if opt.CpuRtPeriod != nil {
		cargs = append(cargs, "--cpu-rt-period="+fmt.Sprint(*opt.CpuRtPeriod))
	}
	if opt.CpuRtRuntime != nil {
		cargs = append(cargs, "--cpu-rt-runtime="+fmt.Sprint(*opt.CpuRtRuntime))
	}
	if opt.CpuShares != nil {
		cargs = append(cargs, "--cpu-shares="+fmt.Sprint(*opt.CpuShares))
	}
	if opt.Cpus != nil {
		cargs = append(cargs, "--cpus="+fmt.Sprint(*opt.Cpus))
	}
	if opt.CpusetCpus != nil {
		cargs = append(cargs, "--cpuset-cpus="+fmt.Sprint(*opt.CpusetCpus))
	}
	if opt.CpusetMems != nil {
		cargs = append(cargs, "--cpuset-mems="+fmt.Sprint(*opt.CpusetMems))
	}
	if opt.KernelMemory != nil {
		cargs = append(cargs, "--kernel-memory="+fmt.Sprint(*opt.KernelMemory))
	}
	if opt.Memory != nil {
		cargs = append(cargs, "--memory="+fmt.Sprint(*opt.Memory))
	}
	if opt.MemoryReservation != nil {
		cargs = append(cargs, "--memory-reservation="+fmt.Sprint(*opt.MemoryReservation))
	}
	if opt.MemorySwap != nil {
		cargs = append(cargs, "--memory-swap="+fmt.Sprint(*opt.MemorySwap))
	}
	if opt.PidsLimit != nil {
		cargs = append(cargs, "--pids-limit="+fmt.Sprint(*opt.PidsLimit))
	}
	if opt.Restart != nil {
		cargs = append(cargs, "--restart="+fmt.Sprint(*opt.Restart))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerVersionOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string

	/*
		Kubernetes config file
	*/
	Kubeconfig *string
}

/*
DockerVersionCmd is wrapper of 'docker version'
------------------------------
version [OPTIONS]
Show the Docker version information
------------------------------
*/
func DockerVersionCmd(opt DockerVersionOption, args []string) *exec.Cmd {
	cargs := []string{"version"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Kubeconfig != nil {
		cargs = append(cargs, "--kubeconfig="+fmt.Sprint(*opt.Kubeconfig))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerVolumeCmd is wrapper of 'docker volume'
------------------------------
volume COMMAND
Manage volumes
------------------------------
*/
func DockerVolumeCmd(args []string) *exec.Cmd {
	cargs := []string{"volume"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerVolumeCreateOption struct {
	/*
		Specify volume driver name
	*/
	Driver *string

	/*
		Set metadata for a volume
	*/
	Label []string

	/*
		Specify volume name
	*/
	Name *string

	/*
		Set driver specific options
	*/
	Opt map[string]string
}

/*
DockerVolumeCreateCmd is wrapper of 'docker volume create'
------------------------------
create [OPTIONS] [VOLUME]
Create a volume
------------------------------
*/
func DockerVolumeCreateCmd(opt DockerVolumeCreateOption, args []string) *exec.Cmd {
	cargs := []string{"volume", "create"}
	if opt.Driver != nil {
		cargs = append(cargs, "--driver="+fmt.Sprint(*opt.Driver))
	}
	if opt.Label != nil {
		for _, str := range opt.Label {
			cargs = append(cargs, "--label")
			cargs = append(cargs, str)
		}
	}
	if opt.Name != nil {
		cargs = append(cargs, "--name="+fmt.Sprint(*opt.Name))
	}
	if opt.Opt != nil {
		for key, val := range opt.Opt {
			cargs = append(cargs, "--opt")
			cargs = append(cargs, key+"="+val)
		}
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerVolumeInspectOption struct {
	/*
		Format the output using the given Go template
	*/
	Format *string
}

/*
DockerVolumeInspectCmd is wrapper of 'docker volume inspect'
------------------------------
inspect [OPTIONS] VOLUME [VOLUME...]
Display detailed information on one or more volumes
------------------------------
*/
func DockerVolumeInspectCmd(opt DockerVolumeInspectOption, args []string) *exec.Cmd {
	cargs := []string{"volume", "inspect"}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerVolumeLsOption struct {
	/*
		Provide filter values (e.g. 'dangling=true')
	*/
	Filter *string

	/*
		Pretty-print volumes using a Go template
	*/
	Format *string

	/*
		Only display volume names
	*/
	Quiet *bool
}

/*
DockerVolumeLsCmd is wrapper of 'docker volume ls'
------------------------------
ls [OPTIONS]
List volumes
------------------------------
*/
func DockerVolumeLsCmd(opt DockerVolumeLsOption, args []string) *exec.Cmd {
	cargs := []string{"volume", "ls"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Format != nil {
		cargs = append(cargs, "--format="+fmt.Sprint(*opt.Format))
	}
	if opt.Quiet != nil {
		cargs = append(cargs, "--quiet="+fmt.Sprint(*opt.Quiet))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerVolumePruneOption struct {
	/*
		Provide filter values (e.g. 'label=<label>')
	*/
	Filter *string

	/*
		Do not prompt for confirmation
	*/
	Force *bool
}

/*
DockerVolumePruneCmd is wrapper of 'docker volume prune'
------------------------------
prune [OPTIONS]
Remove all unused local volumes
------------------------------
*/
func DockerVolumePruneCmd(opt DockerVolumePruneOption, args []string) *exec.Cmd {
	cargs := []string{"volume", "prune"}
	if opt.Filter != nil {
		cargs = append(cargs, "--filter="+fmt.Sprint(*opt.Filter))
	}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

type DockerVolumeRmOption struct {
	/*
		Force the removal of one or more volumes
	*/
	Force *bool
}

/*
DockerVolumeRmCmd is wrapper of 'docker volume rm'
------------------------------
rm [OPTIONS] VOLUME [VOLUME...]
Remove one or more volumes
------------------------------
*/
func DockerVolumeRmCmd(opt DockerVolumeRmOption, args []string) *exec.Cmd {
	cargs := []string{"volume", "rm"}
	if opt.Force != nil {
		cargs = append(cargs, "--force="+fmt.Sprint(*opt.Force))
	}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}

/*
DockerWaitCmd is wrapper of 'docker wait'
------------------------------
wait CONTAINER [CONTAINER...]
Block until one or more containers stop, then print their exit codes
------------------------------
*/
func DockerWaitCmd(args []string) *exec.Cmd {
	cargs := []string{"wait"}
	cargs = append(cargs, args...)
	return exec.Command("docker", cargs...)
}
