package main

import "github.com/junyaU/mimi/cmd"

func main() {
	//command, err := inputparser.NewCommand(os.Args[1:]...)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	//	os.Exit(1)
	//
	//}
	//
	//info, err := pkginfo.New(command.Path)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	//	os.Exit(1)
	//}
	//
	//graph := depgraph.New(info)
	//
	//if command.IsGraph {
	//	graphDrawer, err := output2.NewGraphDrawer(command.MaxDirectDeps, command.MaxIndirectDeps)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	//		os.Exit(1)
	//	}
	//
	//	if err := graphDrawer.Draw(graph.PrintRows()); err != nil {
	//		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	//		os.Exit(1)
	//	}
	//}
	//
	//drawer, err := output2.NewLogDrawer(graph.GetNodes())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	//	os.Exit(1)
	//}
	//
	//if command.IsVerbose {
	//	drawer.Draw()
	//}
	//
	//if command.IsSetMaxDeps() && drawer.ReportExceededDeps(command.MaxDirectDeps, command.MaxIndirectDeps) {
	//	fmt.Fprintf(os.Stderr, "Error: Exceeded max direct or indirect dependencies\n")
	//	os.Exit(1)
	//}
	//

	cmd.Execute()
}
