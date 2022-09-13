# 没

# Step [O1o1o0g9o0]

```go
		var setMsgArgs = func(msg string, args ...interface{}) {
			logg.j.Infow(msg, args)
		}
		logg.Infow(setMsgArgs, "Welcome!", "name", "nihon taro", "weight", 92.6, "x", 3)
```

# Step [O1o1o0g11o__10o2o0]

```go
func (logg *Logger) Infow(setMsgArgs func(string, ...interface{}), msg string, keysAndValues ...interface{}) {
	var sb strings.Builder
	if msg != "" {
		sb.WriteString(msg)
		sb.WriteString(" ")
	}

	for i, v := range keysAndValues {
		if i == 0 {
			// pass
		} else if i%2 == 0 {
			sb.WriteString(" ")
		} else {
			sb.WriteString(":")
		}

		sb.WriteString(fmt.Sprintf("%v", v))
	}
	logg.c.Infof("%s", sb.String())

	setMsgArgs(msg, keysAndValues...)
	// logg.j.Infow(msg, keysAndValues...)
}
```
