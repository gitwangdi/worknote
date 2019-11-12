# Manager and controller

## manager

Contains some informations, like config, scheme, cache and so on.

It can be used to create controller.

When we invoke controller.New():

1. Set controller's fields with manager.
2. Use mgr.Setfield to deal with reconciler.
   > Setfield will check if reconciler has InjectClient(cache,config and so on) func, so it can use the func. I think it is just a way for reconciler to get these info.
3. mgr.Add(controller).
   > This will invoke controller.Start and init controller.

## Controller

### Watch

Use source.Start add event from the source into controller.queue.

source.Start actually start a informer.[link](../informer.md)
