# viper

## Reference

[GitHub](https://github.com/spf13/viper)

## Outline

Viper 是一个用于处理 golang 程序配置信息的包，可以方便地从配置文件、参数（Flag）和环境变量（Env）中读取信息。

## Usage

### Read Configuration File

```golang
func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

    err := viper.ReadInConfig()
	if err != nil {
		return
	}

	fmt.Println(viper.GetString("arg-name"))
}
```

上述代码片段从 `./config.json`（由 `AddConfigPath`、`SetConfigName` 和 `SetConfigType` 三个函数指定）中读取配置信息，然后读取其中名为 `arg-name` 的配置数据。

#### Parse config into a structure

```golang
type Config struct {
	Arg string
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

    err := viper.ReadInConfig()
	if err != nil {
		return
	}

	config := Config{}
	err = viper.Unmarshal(&config)
	fmt.Println(config)
}
```

上述代码片段，将读取到的配置数据解析为 `Config` 结构。

从 Config File、Flag 和 Env 中读取到的数据都可以使用 `viper.Getxxx` 和 `viper.Unmarshal` 两种方式获得配置信息。

### Read Flag

```golang
type Config struct {
	Arg string
}

func main() {
	flagSet := pflag.NewFlagSet("test-viper", pflag.ExitOnError)
	flagSet.String("arg", "arg", "some discription")
	flagSet.Parse(os.Args[1:])

	viper.BindPFlags(flagSet)

	config := Config{}
	viper.Unmarshal(&config)
	fmt.Println(config)
}
```

上述代码片段从接收参数 `--arg`，并解析为 `Config` 格式。

此处有一个值得指明的地方，`Config` 中的字段可以是 `Arg`、`ARg`、`ArG` 或 `ARG`，这些都可以正确读取 `--arg` 和配置文件中 `arg`（后面的 Env 也是如此）。（但是不可以以小写字母开头，这个涉及到 golang 的规则：不可以向其他包中小写字母开头的变量赋值）

#### Read Nest Config

```golang
type Config struct {
	SubConfig struct {
        Arg string
    }
}

func main() {
    flagSet := pflag.NewFlagSet("test-viper", pflag.ExitOnError)
	flagSet.String("subconfig.arg", "arg", "some discription")
	flagSet.Parse(os.Args[1:])

	viper.BindPFlags(flagSet)

	config := Config{}
	viper.Unmarshal(&config)
	fmt.Println(config)
}
```

上述代码片段从接收参数 `--subconfig.arg`，并解析为 `Config` 格式。

#### Array Flag

```golang
flagSet.StringSlice("arg", []string{}, "some discription")
```

参数传入时，格式为 `go run main.go --arg=aa --arg=bb`。

### Read Enviroment

```golang
type Config struct {
	Arg string
}

func main() {
	viper.BindEnv("arg")

	config := Config{}
	viper.Unmarshal(&config)
	fmt.Println(config)
}
```

上述代码片段读取 `ARG` 环境变量，然后解析为 `Config` 格式。

#### Use Env Prefix

```golang
type Config struct {
	Arg string
}

func main() {
	viper.SetEnvPrefix("TEST")
	viper.BindEnv("arg")

	config := Config{}
	viper.Unmarshal(&config)
	fmt.Println(config)
}
```

上述代码片段读取 `TEST_ARG` 环境变量，然后解析为 `Config` 格式。

#### Read Nest Config

```golang
type Config struct {
	SubConfig struct {
        Arg string
    }
}

func main() {
	viper.BindEnv("subconfig.arg")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := Config{}
	viper.Unmarshal(&config)
	fmt.Println(config)
}
```

上述代码片段读取 `SUBCONFIG_ARG` 环境变量，然后解析为 `Config` 格式。

需要注意的是，这里需要使用 `SetEnvKeyReplacer` 将嵌套式配置中的嵌套结构的 `.` 转为 `_`，因为环境变量中不应有 `.`。

#### Automatic Env

```golang
func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

    err := viper.ReadInConfig()
	if err != nil {
		return
	}

    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	fmt.Println(viper.GetString("arg-name"))
}
```

使用 viper.Getxxx 读取配置数据时，可以使用 `AutomaticEnv` 函数。

使用 viper.Unmarshal 读取配置数据时，使用 `AutomaticEnv` 读取的环境变量应该是 Config File、Flag 或 BindEnv 等方式注册过的变量，否则无法正确的解析到 `Config` 结构中（即使 `Config` 中有这个变量的定义，但是在 Config File 中没有该变量，则不能正确读取环境变量）。

#### Array Env

环境变量设置方式为 `ARG="env1 env2 env3"`，可以使用 `GetStringSlice(GetIntSlice ...)` 和 `Unmarshal` 方式解析。

### Combine all

```golang
type Config struct {
	DB struct {
		Host string
        Port int
	}
}

func main() {
	flagSet := pflag.NewFlagSet("db", pflag.ExitOnError)
	flagSet.String("db.host", "localhost", "database host")
    flagSet.Int("db.port", 80, "database port")
	flagSet.Parse(os.Args[1:])
	viper.BindPFlags(flagSet)

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("json")
    err := viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.SetEnvPrefix("CONFIG")
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	config := Config{}
	viper.Unmarshal(&config)
	fmt.Println(config)
}
```

上述代码片段可以从同时从配置文件、参数和环境变量中读取配置数据。同一项配置的覆盖关系为：参数 > 环境变量 > 配置文件 > 参数默认值。
