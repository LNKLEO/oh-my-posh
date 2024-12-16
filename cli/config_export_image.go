package cli

import (
	"fmt"

	"github.com/LNKLEO/OMP/config"
	"github.com/LNKLEO/OMP/image"
	"github.com/LNKLEO/OMP/prompt"
	"github.com/LNKLEO/OMP/runtime"
	"github.com/LNKLEO/OMP/shell"
	"github.com/LNKLEO/OMP/template"
	"github.com/LNKLEO/OMP/terminal"

	"github.com/spf13/cobra"
)

var (
	// cursorPadding int
	// rPromptOffset int
	author      string
	bgColor     string
	outputImage string
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Export your config to an image",
	Long: `Export your config to an image.

You can tweak the output by using additional flags:

- author: displays the author below the prompt
- cursor-padding: the padding of the prompt cursor
- rprompt-offset: the offset of the right prompt
- background-color: the background color of the image

Example usage:

> oh-my-posh config export image --config ~/myconfig.omp.json

Exports the config to an image file called myconfig.png in the current working directory.

> oh-my-posh config export image --config ~/myconfig.omp.json --output ~/mytheme.png

Exports the config to an image file ~/mytheme.png.

> oh-my-posh config export image --config ~/myconfig.omp.json --author "John Doe"

Exports the config to an image file using customized output options.`,
	Args: cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		env := &runtime.Terminal{
			CmdFlags: &runtime.Flags{
				Config:        configFlag,
				Shell:         shell.GENERIC,
				TerminalWidth: 150,
			},
		}

		env.Init()
		defer env.Close()

		template.Init(env)

		cfg := config.Load(env)

		// set sane defaults for things we don't print
		cfg.ConsoleTitleTemplate = ""
		cfg.PWD = ""

		// add variables to the environment
		env.Var = cfg.Var

		terminal.Init(shell.GENERIC)
		terminal.BackgroundColor = cfg.TerminalBackground.ResolveTemplate()
		terminal.Colors = cfg.MakeColors()

		eng := &prompt.Engine{
			Config: cfg,
			Env:    env,
		}

		primaryPrompt := eng.Primary()

		imageCreator := &image.Renderer{
			AnsiString: primaryPrompt,
			Author:     author,
			BgColor:    bgColor,
		}

		if outputImage != "" {
			imageCreator.Path = cleanOutputPath(outputImage, env)
		}

		err := imageCreator.Init(env)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		err = imageCreator.SavePNG()
		if err != nil {
			fmt.Print(err.Error())
		}
	},
}

func init() {
	imageCmd.Flags().StringVar(&author, "author", "", "config author")
	imageCmd.Flags().StringVar(&bgColor, "background-color", "", "image background color")
	imageCmd.Flags().StringVarP(&outputImage, "output", "o", "", "image file (.png) to export to")
	exportCmd.AddCommand(imageCmd)
}
