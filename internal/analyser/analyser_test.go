	"context"
func (a *mockAnalyser) NewExecuter(_ context.Context, _ string) (Executer, error) {
func (a *mockAnalyser) Execute(_ context.Context, args []string) (out []byte, err error) {
func (a *mockAnalyser) Stop(_ context.Context) error {
	issues, err := Analyse(context.Background(), analyser, tools, cfg)
	issues, err := Analyse(context.Background(), analyser, tools, cfg)
	_, err := Analyse(context.Background(), analyser, nil, cfg)