package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var selected bool
var emaillike bool

func init() {
	f := rootCmd.Flags()
	f.BoolVarP(&selected, "select", "s", false, "Select payload and just commit selected files")
	f.BoolVarP(&emaillike, "emaillike", "e", false, "Write an email like (longer) commit message. Uses your default $EDITOR")
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "gg",
	Short: "Add, Commit, Push be happy with gg",
	Long: `Shell promt to add, commit, and push your amazing stuff to 
	git version controll without typos and faster. I swear. `,

	Run: func(cmd *cobra.Command, args []string) {
		if !isRepo() {
			log.Fatal("fatal: not a git repository. Run git init..")
			return
		}

		if !isDiverged() {
			if selected == true {
				payload := []string{}
				prompt := &survey.MultiSelect{
					Message:  "Select payload wisely! (What you select here shall be committed!)",
					Options:  getOptions(),
					PageSize: 15,
				}
				survey.AskOne(prompt, &payload)
				addSelected(payload)
				status := commitStuff()
				fmt.Println(status)
				return
			}

			addAll()
			status := commitStuff()
			fmt.Println(status)
			return
		}

		fmt.Println("Please resolve/pull first. Your current branch is behind!")
		return
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("could not execute cobra command %v \n", err)
	}
}

func commitStuff() string {
	var message string
	if emaillike == true {
		prompt := &survey.Editor{
			Message:  "Please provide a good commit message in simple present! (You will be redirected to vim )",
			FileName: "commit_message.txt",
			Editor:   os.Getenv("EDITOR"),
		}
		survey.AskOne(prompt, &message)
		commit(message)
		pushOut, _ := push()
		fmt.Println(string(pushOut))
		return status()
	}
	prompt := &survey.Multiline{Message: "Please provide a good commit message in simple present!"}
	survey.AskOne(prompt, &message)
	commit(message)
	pushOut, _ := push()
	fmt.Println(string(pushOut))

	return status()
}
