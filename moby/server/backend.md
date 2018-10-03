# backend

## struct Backend 

```go
type ImageComponent interface {
	SquashImage(from string, to string) (string, error)
	TagImageWithReference(image.ID, reference.Named) error
}

// Builder defines interface for running a build
type Builder interface {
	Build(context.Context, backend.BuildConfig) (*builder.Result, error)
}

type Backend struct {
    builder          Builder
    fsCache          *fscache.FSCache
    imageComponent   ImageComponent
    buildkit         *buildkit.Builder
}
```

用来处理一个叫build的东西，建立起build source，然后从中建立image  


## struct Tagger

```go
type Tagger struct {
	imageComponent ImageComponent
	stdout         io.Writer
	repoAndTags    []reference.Named
}
```

NewTagger() TagImage()  
标记image