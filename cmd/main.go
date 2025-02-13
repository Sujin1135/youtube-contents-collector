package main

import (
	"channel-contents-collector/pkg/collector"
	"channel-contents-collector/pkg/external"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>

	dataApi := external.NewDataAPI()
	channelCollector := collector.NewChannelCollector(dataApi)
	contentCollector := collector.NewContentCollector(dataApi)

	// TODO: collect
	fmt.Println(channelCollector.Collect())
	fmt.Println(contentCollector.Collect())
}
