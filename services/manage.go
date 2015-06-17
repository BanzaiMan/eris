package services

import (
  "fmt"
  "regexp"
  // "strings"

  "github.com/eris-ltd/eris-cli/util"

  "github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/spf13/cobra"
)

// install
func Install(cmd *cobra.Command, args []string) {

}

func Configure(cmd *cobra.Command, args []string) {

}

func Inspect(cmd *cobra.Command, args []string) {
  imgs, _ := util.DockerClient.ListImages(docker.ListImagesOptions{All: false})
  for _, img := range imgs {
    fmt.Println("ID: ", img.ID)
    fmt.Println("RepoTags: ", img.RepoTags)
    fmt.Println("Created: ", img.Created)
    fmt.Println("Size: ", img.Size)
    fmt.Println("VirtualSize: ", img.VirtualSize)
    fmt.Println("ParentId: ", img.ParentID)
  }
}

// Updates an installed service, or installs it if it has not been installed.
func Update(cmd *cobra.Command, args []string) {

}

// list known
func ListKnown() {

}

func ListRunning() {
  services := ListRunningRaw()
  for _, s := range services {
    fmt.Println(s)
  }
}

func ListInstalled() {

}

func Rm(cmd *cobra.Command, args []string) {

}

func ListRunningRaw() []string {
  services := []string{}
  r := regexp.MustCompile(`\/eris_service_(.+)_\S+?_\d`)

  contns, _ := util.DockerClient.ListContainers(docker.ListContainersOptions{All: false})
  for _, con := range contns {
    for _, c := range con.Names {
      match := r.FindAllStringSubmatch(c, 1)
      if len(match) != 0 {
        services = append(services, r.FindAllStringSubmatch(c, 1)[0][1])
      }
    }
  }

  return services
}

func IsServiceRunning(service *util.Service) bool {
  running := ListRunningRaw()
  if len(running) != 0 {
    for _, srv := range running {
      if srv == service.Name {
        return true
      }
    }
  }
  return false
}