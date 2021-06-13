local TILE_EMPTY = 0
local TILE_BLOCK = 1

function LoadLevel(level)
    map = ""
    size = {w=15,h=16}


    if level == 1 then
        map =
        ".#.#.#.#.#.#.#."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."..
        "..............."
    end
    _CreateLevel(size.w, size.h)
    for y = 1, size.h-1 do
        for x = 1, size.w do
            c = string.sub(map, ((y-1)* size.w + x), ((y-1) * size.w +x) )
            if      c == '.' then _SetTile(x-1, y-1, TILE_EMPTY)
            elseif  c == '#' then _SetTile(x-1, y-1, TILE_BLOCK)
            end
        end
    end
    
end
