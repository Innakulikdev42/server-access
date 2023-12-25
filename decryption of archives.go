func TestCreate(t *testing.T) {
	preTestValidate(t)

	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx
	cli.ForceSkipConfirm = true

  	t.Run("Recipe", func(t *testing.T) {
		args := []string{"NewUser0", goodPLR}
		cmd := DevCreate()
		_, err := clitestutil.ExecTestCLICmd(ctx, cmd, args)
		assert.Nil(t, err)
	})

  func DevValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate [path]",}}}
