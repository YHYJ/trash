- [ ] restore 子命令：
  - [ ] 恢复回收站中的文件（可用参数为 --help） (2023-11-26 14:02)
  - [ ] 子命令执行后输出结果格式是：文件编号（0是全部文件） 删除日期 删除时间 文件原路径 (2023-11-26 14:04)
  - [ ] 然后是询问想要恢复哪个文件的提示 (2023-11-26 14:04)
  - [ ] 恢复文件前检测文件原路径是否已有了同名文件并提示是否覆盖 (2023-11-26 14:05)
  - [ ] 回收站本来就为空时给予提示 (2023-11-26 14:09)
  - [ ] 未检测到回收站文件夹则创建（权限=755、属主/属组=当前用户） (2023-11-26 13:53)
- [ ] purge 子命令：
  - [ ] 清空回收站（可用参数为 --help） (2023-11-26 14:07)
  - [ ] 需要二次确认 (2023-11-26 14:07)
  - [ ] 附带清空记录文件（.trashinfo文件） (2023-11-26 14:10)
- [X] put 子命令：
  - [X] 将文件移动到回收站（可用参数为 --help 或要删除的文件名） (2023-11-26 13:50)
  - [X] 指定文件未找到时提示没有该文件 (2023-11-26 13:50)
  - [X] 记录文件的原地址和删除日期（.trashinfo文件） (2023-11-26 13:51)
  - [X] 未检测到回收站文件夹则创建（权限=755、属主/属组=当前用户） (2023-11-26 13:53)
- [X] list 子命令：
  - [X] 列出回收站中的文件（可用参数为 --help） (2023-11-26 13:54)
  - [X] 子命令执行后输出结果格式是：删除日期 删除时间 文件原路径 (2023-11-26 13:56)
  - [X] 未检测到回收站文件夹则创建（权限=755、属主/属组=当前用户） (2023-11-26 13:53)