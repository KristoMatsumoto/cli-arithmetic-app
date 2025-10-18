package app

type PipelineOptions struct {
	InputPath      string
	OutputPath     string
	Format         string
	ProcessorType  string
	Version        string
	TransformChain []string
}
