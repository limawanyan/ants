package ants

import "time"

// Option represents the optional function.
type Option func(opts *Options)

func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

// Options 实例化协程池的相关参数
// Options contains all options which will be applied when instantiating an ants pool.
type Options struct {
	// 定期清除超时未执行任务时间间隔
	// ExpiryDuration is a period for the scavenger goroutine to clean up those expired workers,
	// the scavenger scans all workers every `ExpiryDuration` and clean up those workers that haven't been
	// used for more than `ExpiryDuration`.
	ExpiryDuration time.Duration

	// 初始化协程池时是否进行内存初始化
	// PreAlloc indicates whether to make memory pre-allocation when initializing Pool.
	PreAlloc bool

	// pool.Submit最大协程阻塞数
	// 0表示没有限制
	// Max number of goroutine blocking on pool.Submit.
	// 0 (default value) means no such limit.
	MaxBlockingTasks int

	// 为ture时,Pool.Submit永远不会被阻止
	// Pool.Submit不能一次性完成时,将返回ErrPoolOverload
	// 为true时MaxBlockingTasks无效
	// When Nonblocking is true, Pool.Submit will never be blocked.
	// ErrPoolOverload will be returned when Pool.Submit cannot be done at once.
	// When Nonblocking is true, MaxBlockingTasks is inoperative.
	Nonblocking bool

	// PanicHandler 用于处理来自每个工作 goroutine 的恐慌。
	// 如果为零，则将再次从工作 goroutine 中抛出恐慌。
	// PanicHandler is used to handle panics from each worker goroutine.
	// if nil, panics will be thrown out again from worker goroutines.
	PanicHandler func(interface{})

	// 用自定义日志记录器,如果未设置则使用日志包中的默认标准记录器。
	// Logger is the customized logger for logging info, if it is not set,
	// default standard logger from log package is used.
	Logger Logger
}

// WithOptions 设置配置选项
// WithOptions accepts the whole options config.
func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

// WithExpiryDuration 设置清理 goroutine 的间隔时间
// WithExpiryDuration sets up the interval time of cleaning up goroutines.
func WithExpiryDuration(expiryDuration time.Duration) Option {
	return func(opts *Options) {
		opts.ExpiryDuration = expiryDuration
	}
}

// WithPreAlloc 设置是否进行内存初始化
// WithPreAlloc indicates whether it should malloc for workers.
func WithPreAlloc(preAlloc bool) Option {
	return func(opts *Options) {
		opts.PreAlloc = preAlloc
	}
}

// WithMaxBlockingTasks 设置达到池容量时阻塞的最大 goroutine 数
// WithMaxBlockingTasks sets up the maximum number of goroutines that are blocked when it reaches the capacity of pool.
func WithMaxBlockingTasks(maxBlockingTasks int) Option {
	return func(opts *Options) {
		opts.MaxBlockingTasks = maxBlockingTasks
	}
}

// WithNonblocking 当没有可用的工作时，池将返回 nil
// WithNonblocking indicates that pool will return nil when there is no available workers.
func WithNonblocking(nonblocking bool) Option {
	return func(opts *Options) {
		opts.Nonblocking = nonblocking
	}
}

// WithPanicHandler 设置Panic处理方法
// WithPanicHandler sets up panic handler.
func WithPanicHandler(panicHandler func(interface{})) Option {
	return func(opts *Options) {
		opts.PanicHandler = panicHandler
	}
}

// WithLogger 设置日志记录器
// WithLogger sets up a customized logger.
func WithLogger(logger Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}
