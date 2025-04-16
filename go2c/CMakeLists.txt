cmake_minimum_required(VERSION 3.10)

project(go2c LANGUAGES CXX)

# 设置C++标准
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# 添加编译选项，类似于Makefile中的CXXFLAGS
add_compile_options(-Wall)

# 添加可执行文件
add_executable(
    go2c
    main.cpp
)

# 添加库搜索路径（当前目录）
target_link_directories(
    go2c
    PRIVATE
    ${CMAKE_CURRENT_SOURCE_DIR}
)

# 添加头文件搜索路径
target_include_directories(
    go2c
    PRIVATE
    ${CMAKE_CURRENT_SOURCE_DIR}
)

# 链接libshare库
target_link_libraries(
    go2c
    PRIVATE
    libshare
)

# 复制DLL到输出目录（确保运行时能找到DLL）
add_custom_command(
    TARGET go2c
    POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CMAKE_CURRENT_SOURCE_DIR}/libshare.dll"
        "$<TARGET_FILE_DIR:go2c>/libshare.dll"
)
