import './wasm/wasm_exec.js';
import * as View from './View.mjs';

for (let x in View) {
    window[x] = View[x];
}

const go = new Go();
console.log("Loading...");
WebAssembly.instantiateStreaming(fetch('/wasm/main.wasm'), go.importObject).then((result) => {
    console.log("Running...");
    go.run(result.instance);
});
