<template>
  <div ref="terminal" style="width: 1280px;height: 1000px;"></div>
</template>

<script>
import "xterm/dist/xterm.css";
import "xterm/dist/addons/fullscreen/fullscreen.css";

import { Terminal } from "xterm";
import * as fit from "xterm/lib/addons/fit/fit";
import * as fullscreen from "xterm/lib/addons/fullscreen/fullscreen";
import * as webLinks from "xterm/lib/addons/webLinks/webLinks";
import * as attach from "xterm/lib/addons/attach/attach";

export default {
  name: "Index",
  created() {
    Terminal.applyAddon(attach);
    Terminal.applyAddon(fit);
    Terminal.applyAddon(fullscreen);
    Terminal.applyAddon(webLinks);

    const terminal = new Terminal();
    const ws = new WebSocket("ws://192.168.17.187:8000/terminal?container=a54c54556b45");
    ws.onclose = function() {
      console.log("服务器关闭了连接");
    };
    terminal.open(this.$refs.terminal);
    terminal.fit();
    terminal.toggleFullScreen();
    terminal.webLinksInit();
    terminal.attach(ws);
  }
};
</script>
