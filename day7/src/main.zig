const std = @import("std");

pub fn main() anyerror!void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const alloc = gpa.allocator();

    const f = try std.fs.cwd().openFile("input", .{ .read = true });
    const input = try f.readToEndAlloc(alloc, 100 * 1_000_000);
    defer alloc.free(input);

    var it = std.mem.tokenize(u8, input, ", \n");

    var positions = std.ArrayList(u32).init(alloc);
    defer positions.deinit();
    while (it.next()) |tok| {
        try positions.append(try std.fmt.parseInt(u32, tok, 10));
    }

    // sort to get median
    std.sort.sort(u32, positions.items, {}, comptime std.sort.asc(u32));

    const index: usize = positions.items.len / 2;
    const target_position: u32 = positions.items[index];

    // part one calculate fuel
    var fuel_spent: i64 = 0;
    for (positions.items) |pos| {
        const fuel: i64 = @as(i64, target_position) - @as(i64, pos);
        if (fuel < 0) {
            fuel_spent += -fuel;
        } else {
            fuel_spent += fuel;
        }
    }
    std.debug.print("part one: {}\n", .{fuel_spent});

    // part two
    //
    // sledgehammer approach (calculate each consumption)
    var consumptions = std.ArrayList(i64).init(alloc);
    defer consumptions.deinit();
    for (positions.items) |item| {
        try consumptions.append(calc_fuel(item, positions.items));
    }
    std.sort.sort(i64, consumptions.items, {}, comptime std.sort.asc(i64));

    std.debug.print("part two: {}\n", .{consumptions.items[0]});

    var total: u64 = 0;
    for (positions.items) |pos| total += pos;

    const mean: u64 = total / positions.items.len;
    std.debug.print("part two: {} (using mean)\n", .{calc_fuel(@intCast(u32, mean), positions.items)});
}

// fuel callculation for part two
pub fn calc_fuel(target: u32, positions: []u32) i64 {
    var fuel_spent: i64 = 0;
    for (positions) |pos| {
        var fuel: i64 = @as(i64, target) - @as(i64, pos);
        if (fuel < 0) {
            fuel = -fuel;
        }
        fuel_spent += @divExact((fuel * (fuel + 1)), 2);
    }
    return fuel_spent;
}
