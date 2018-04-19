import React from 'react';
import { render } from 'react-dom';
import request from 'superagent';


class App extends React.Component {
    render() {
        return "HOWDY";
    }
}
render(<App />, document.getElementById("app"));