- 结构体标签
使用 `toml:"field_name"` 标签映射 TOML 键名（标签可省略，但需保持字段名与 TOML 键名一致）。

- 嵌套结构
TOML 的 [section] 对应 Go 的嵌套结构体。

- 数组处理
TOML 的 [[array]] 对应 Go 的切片（[]Type）。

- 数据类型映射：

TOML 字符串 → Go string

TOML 整数 → Go int

TOML 布尔值 → Go bool

TOML 数组 → Go 切片（[]Type）

TOML 表（Table）→ Go 结构体