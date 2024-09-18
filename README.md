# Go Config Loader 

This package is a really lightweight environment-type config loader that enhance the normal use case of `os.Getenv("ENV_KEY")`. This package will help us get the expected datatype.

This package can load environment data that is defined from OS level, hardcoded map[string]string, or .env file easily. 

### How to use
- Run `go get github.com/tianrosandhy/goconfigloader`
- Init new configloader : `cfg := goconfigloader.NewConfigLoader()`
- Get value : 
    - `cfg.GetString("YOUR_CONFIG_KEY")`
    - `cfg.GetInt("YOUR_CONFIG_KEY")`
    - `cfg.GetFloat64("YOUR_CONFIG_KEY")`
    - `cfg.GetBool("YOUR_CONFIG_KEY")`

### Use cases
1. Define default config from map[string]string
Optinally, you can define a `map[string]string` with pairs of default key=>value. 
```
defaultConfig := map[string]string{
    "API_URL": "http://localhost:9090",
    "PORT": "9090",
    "IS_API_ENABLED": "true",
}
cfg := goconfigloader.NewConfigLoader(defaultConfig)
//
cfg.GetBool("IS_API_ENABLED") // return TRUE
```

2. Define default config from .env file
If you have `.env` file stored in your project path, this package will try to load that file too 
```
API_URL="http://localhost:9000"
PORT=9090
IS_API_ENABLED=true
```

```
cfg := goconfigloader.NewConfigLoader(defaultConfig)
//
cfg.GetInt("PORT") // return 9090
```

3. Define environment from OS level 
By default this package will call the `os.Getenv("KEY")` first, so any environment that you define in OS level will be handled like usual.