import "./style";
import { Component } from "preact";

const addRoute = route =>
  fetch("localhost:3000/api" + route, { method: "POST" });

export default class App extends Component {
  render() {
    return (
      <div>
        <h1>Hello, World!</h1>
        <button onClick={addRoute} />
      </div>
    );
  }
}
