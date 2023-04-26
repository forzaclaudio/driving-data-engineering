# Driving Data Engineering

Enable a pipeline for the collection of real-time F1-simulation telemetry.

## Methodology

1. Start a session in Azure and verify that *`f1sim`* storage is active.
2. Plug in the HDMI from the simulator to the video capture device in the controller procesor.
3. Plug in the HDMI output from the controller processor to the simulator projector/monitor.
4. Start the f1 simulator.
5. Start the OBS studio and add the video capture device as video input, ensure that the webserver is active.
6. Modify the storage settings of the muos server and ensure that the *container* matches the session, e.g. bahrain
7. Start the *`muos`* server.
8. Modify session.toml file to ensure that the driver and gender fields are correct.
9. Start the pipeline.
