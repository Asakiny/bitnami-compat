/*
Copyright © 2022 ZCube <zcubekr@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
  "fmt"
  "io/ioutil"
  "log"
  //"os"
  //"os/exec"

  "github.com/bmatcuk/doublestar/v4"
  // "github.com/kyokomi/emoji/v2"
  "github.com/spf13/cobra"
  "golang.org/x/exp/slices"
  "gopkg.in/yaml.v3"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
  Use:   "generate",
  Short: "generate dockerfile for arm64",
  Long:  `generate dockerfile for arm64`,
  Run: func(cmd *cobra.Command, args []string) {
    // emoji.Println("Clean bitnami containers")
    /**{
      args := []string{"submodule", "update", "--init", "--recursive", "--checkout", "-f"}
      cmd := exec.Command("git", args...)
      cmd.Stdout = os.Stdout
      cmd.Stderr = os.Stderr
      if err := cmd.Run(); err != nil {
        log.Panic(err)
      }
    }
    **/
    buf, err := ioutil.ReadFile("config.yaml")
    if err != nil {
      log.Panic(err)
    }

    p := &Config{}
    err = yaml.Unmarshal(buf, p)
    if err != nil {
      log.Fatalf("Unmarshal: %v", err)
    }

    var dockerfiles []string
    if len(app) > 0 {
      dockerfiles, err = doublestar.FilepathGlob(fmt.Sprintf("containers/bitnami/%v/**/Dockerfile", app))
    } else {
      dockerfiles, err = doublestar.FilepathGlob(fmt.Sprintf("containers/bitnami/*/**/Dockerfile"))
    }
    if err != nil {
      log.Panic(err)
    }

    for i := range dockerfiles {
      if appInfo, err := InspectDockerfile(dockerfiles[i]); err == nil {
        if slices.Contains(p.Ignores, appInfo.Name) {
          continue
        }
        PatchDockerfile(appInfo)
      } else {
        log.Panic(err)
      }
    }
  },
}

func init() {
  rootCmd.AddCommand(generateCmd)
  generateCmd.PersistentFlags().StringVar(&app, "app", "", "app")
}
