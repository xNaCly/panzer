package keywords

// TODO: add source <file>
// TODO: add help [topic] - should print general info without 'topic',
// otherwise specific information for 'topic', such as prompt placeholders,
// config lookup algorithm etc

// maps keywords to their functions
var KEYWORD_MAP = map[string]func(args ...string){
	"env":   Env,
	"set":   Set,
	"cd":    Cd,
	"alias": Alias,
}
