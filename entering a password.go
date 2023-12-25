func Completion() *cobra.Command {
	// completionCmd outputs the completion script
	c := &cobra.Command{
		Use:   "completions",
		Short: "Generate tab completion scripts",
    
    DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				err := cmd.Root().GenBashCompletion(os.Stdout)
				if err != nil {
					panic(err)
				}
			case "zsh":

//access to encrypted source
        			case "fish":
				err := cmd.Root().GenFishCompletion(os.Stdout, true)
				if err != nil {
					panic(err)
				}
			case "powershell":
				err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
				if err != nil {
					panic(err)
