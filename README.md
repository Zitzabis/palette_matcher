# Minecraft Skin Palette Matcher API

## Purpose
This API should be used to check if a Minecraft skin matches against a defined palette. This is highly beneficial for palette contests on PMC (official/unofficial).

## Endpoints

### Match: 
`/api/palette_matcher/match`

The `match` endpoint is the primary API endpoint. It's return data will appear as a boolean:
```json
{"match":true}
```


## Example JSON POST data
Below is an example of JSON data that can be sent via POST.
```json
{
"Hexs":[ "572F4E", "573E69", "4B5F90", "3985B1", "39A2B8", "55C4C1" ],
"URL":"https://www.test.com/path/to/skin.png"
}   
```
- Any number of hex codes can be submitted.
- The URL must be a direct link.


## Speed
On average, each request takes **46ms**.

Connection speed is not a factor in these tests. This is due to the fact that locations can vary.

