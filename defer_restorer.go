package errf

// DeferRestorer is helper interface to update global config in a
// scope of a single function, typically useful for testing or
// for main function configuration.
//
// Example:
//   func Test1(t *testing.Test) {
//     defer SetGlobalConfig(&Config{ param1: "value1", param2: "value2" }).ThenRestore()
//
//     // ...
//   }
type DeferRestorer interface {

	// ThenRestore restores previous global config.
	// See DeferRestorer comment for example.
	ThenRestore()
}
