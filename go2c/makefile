CC=gcc
CXX=g++

# 编译标志
CXXFLAGS=-Wall -std=c++11
# 链接标志，添加当前目录到库搜索路径
LDFLAGS=-L.

# 默认目标
all: main

# 构建主程序
main: main.cpp
	$(CXX) $(CXXFLAGS) -o $@ $< $(LDFLAGS) -llibshare

# 直接链接DLL的替代方法
main_direct: main.cpp
	$(CXX) $(CXXFLAGS) -o $@ $< -L. libshare.dll

# 清理生成的文件
clean:
	del main.exe main_direct.exe *.o

# 运行程序(设置PATH以找到DLL)
run: main
	main

.PHONY: all clean run
