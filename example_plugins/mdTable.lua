function process(input)
	local lines = splitStr(input, "\n")
	local grid, counts = splitCols(lines)
	local paddedGrid = padCells(grid, counts)

	local output = printTable(paddedGrid)
	return output
end

function printTable(grid)
	local output = {}
	for r=0, #grid, 1 do
		output[r] = table.concat(grid[r], "|")
	end

	return "|" .. table.concat(output, "|\n|") .. "|"
end

function padCells(grid, counts)
	local output = {}
	local oi = 1
	output[0] = {""}
	for r=0, #grid, 1 do
		if r == 1 then
			output[oi] = {}
			for c=0, #counts, 1 do
				output[oi][c] = string.rep("-", counts[c]+2)
			end
			oi = oi + 1
		end
		output[oi] = {}
		for c=0, #counts, 1 do
			if counts[c] == nil then
				counts[c] = 0
			end
			if grid[r] == nil then
				grid[r] = {}
			end
			if c > #grid[r] then
				grid[r][c] = ""
			end
			output[oi][c] = " " .. string.format("%-"..counts[c].."s", grid[r][c]) .. " "
		end
		oi = oi + 1
	end

	return output
end

function trim(s)
	return (s:gsub("^%s*(.-)%s*$", "%1"))
end

function splitCols(lines)
	local output = {}
	local counts = {}
	local oi = 0
	for r = 0, #lines, 1 do
		if lines[r] == nil or string.match(lines[r], "^[- |]*$") then
			goto continue
		end

		output[oi] = {}
		local cols = splitStr(lines[r], "|")

		for c=0, #cols, 1 do
			if cols[c] ~= nil then
				local cell = trim(cols[c])
				local cellCount = #cell
				if #counts < c or counts[c] == nil or counts[c] < cellCount then
					counts[c] = cellCount
				end
				output[oi][c] = cell
			end
		end

		oi = oi + 1

		::continue::
	end

	return output, counts
end

function splitStr(input, sep)
	if input == nil then
  		return {}
	end
	if sep == nil then
		sep = "|"
	end
	local output = {}
	for str in string.gmatch(input, "([^"..sep.."]+)") do
		table.insert(output, str)
	end
	return output
end
