# vut - Very Useful Tool


A simple tool to manage volume, mic, brightness etc,
and anything that can be thrown in this general framework.



# Roadmap

* refine the design using CLI
- [ ] brightnessctl
- [ ] nmcli c
- [ ] pactl
- [ ] pwctl
- [ ] wlrandr

* make GUI using fyne

## CLI Usage

###  vut
```
 > vut
  Device    |         Name           | Value |  Range
brightness  |  intel_backlight       |   960 |  0 .. 96000
brightness  |  phy0-led              |     1 |  0 .. 1
brightness  |  input3::numlock       |     0 |  0 .. 1
brightness  |  input3::capslock      |     0 |  0 .. 1
brightness  |  input3::scrolllock    |     0 |  0 .. 1
brightness  |  dell::kbd_backlight   |     0 |  0 .. 2
brightness  |  platform::micmute     |     0 |  0 .. 1
```


### vut <pat>
```
> vut intel
  Device    |         Name           | Value |  Range
brightness  |  intel_backlight       |   960 |  0 .. 96000
```

### vut <pat> <val>
```
> vut intel 100%
  Device    |         Name           | Value |  Range
brightness  |  intel_backlight       | 96000 |  0 .. 96000
```
