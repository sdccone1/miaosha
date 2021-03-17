/**
 * @Author:David Ma
 * @Date:2021-02-01
 */

package logService

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
 * 获取日志对象
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	//开启文件和行号
	caller := zap.AddCaller()
	//开启开发者模式，也就是stacktrace
	development := zap.Development()
	//往log中追加额外的信息(要求是<k,v>类型)
	opts := zap.Fields(zap.String("service", serviceName))
	logger := zap.New(core, caller, development, opts)
	return logger
}

/**
 * 构建zapcore
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {

	//利用lumberjack注册一个hook也就是一个可以进行分割的logger
	hook := lumberjack.Logger{
		Filename:   filePath,   //文件路径
		MaxSize:    maxSize,    //每个日志文件的大小，单位：M
		MaxBackups: maxBackups, //每个日志文件的备份的个数
		MaxAge:     maxAge,     //每个日志文件的保留时长,单位：天
		Compress:   compress,
	}

	//自定义一个编码器encoder
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "name",
		CallerKey:        "location", //在日志中lines这个<k,v>对的value表示的即为触发将msg写在log的代码所在的文件及行号
		FunctionKey:      "func",
		StacktraceKey:    "stacktrace",                   //development mode下生效
		LineEnding:       zapcore.DefaultLineEnding,      //日志每一行的结尾默认为换行符'\n'
		EncodeLevel:      zapcore.LowercaseLevelEncoder,  //小写的编码器
		EncodeTime:       zapcore.ISO8601TimeEncoder,     //ISO8601 UTC 时间格式
		EncodeDuration:   zapcore.SecondsDurationEncoder, //编码的所需时间
		EncodeCaller:     zapcore.FullCallerEncoder,      //全路径编码器
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: " ",
	})

	core := zapcore.NewCore(encoder, //编码器
		zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(&hook)), //打印到控制台和文件
		level, //日志级别
	)
	return core
}
