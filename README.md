# comProxy
A go service listening to comunication bettwen ground control and drone by com port, extracting flight data from it and making it avaliable
with REST endpoints.
## Available endpoints:
### Gps data
```GET /gps```
```json
{  
    "TimeBootMs":0,
    "Lat":0,
    "Lon":0,
    "Alt":0,
    "RelativeAlt":0,
    "Vx":0,
    "Vy":0,
    "Vz":0,
    "Hdg":0
}
```
```
TimeBootMs  - (uint32) Timestamp (milliseconds since system boot)
Lat         - (int32)  Latitude, expressed as * 1E7
Lon         - (int32)  Longitude, expressed as * 1E7
Alt         - (int32)  Altitude in meters, expressed as * 1000 (millimeters), above MSL
RelativeAlt - (int32)  Altitude above ground in meters, expressed as * 1000 (millimeters)
Vx          - (int16)  Ground X Speed (Latitude), expressed as m/s * 100
Vy          - (int16)  Ground Y Speed (Longitude), expressed as m/s * 100
Vz          - (int16)  Ground Z Speed (Altitude), expressed as m/s * 100
Hdg         - (uint16) Compass heading in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
```
### Accelerometer data
```GET /attitude```
```json
{  
    "TimeBootMs":0,
    "Roll":0,
    "Pitch":0,
    "Yaw":0,
    "Rollspeed":0,
    "Pitchspeed":0,
    "Yawspeed":0
}
```
```
TimeBootMs - (uint32)  Timestamp (milliseconds since system boot)
Roll       - (float32) Roll angle (rad, -pi..+pi)
Pitch      - (float32) Pitch angle (rad, -pi..+pi)
Yaw        - (float32) Yaw angle (rad, -pi..+pi)
Rollspeed  - (float32) Roll angular speed (rad/s)
Pitchspeed - (float32) Pitch angular speed (rad/s)
Yawspeed   - (float32) Yaw angular speed (rad/s)
```
