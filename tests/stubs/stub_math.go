package stubs

type StubMathRepository struct{}

func (s *StubMathRepository) Add(a, b int) int {
	// Retorna um valor fixo independentemente dos inputs
	return 42
}
