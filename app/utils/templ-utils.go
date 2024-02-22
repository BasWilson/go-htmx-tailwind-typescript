package utils

func CreateClassName(classNames ...string) string {
	className := ""
	for _, name := range classNames {
		if name != "" {
			className += " " + name
		}
	}
	return className
}
