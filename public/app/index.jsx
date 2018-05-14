import React from 'react';
import { render } from 'react-dom';
import request from 'superagent';


class App extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            poos: []
        }

        this.getPoos = this.getPoos.bind(this)
        this.flush = this.flush.bind(this)
    }

    componentDidMount() {
        this.getPoos()
    }

    getPoos() {
        request
            .get('/api/poos')
            .set('accept', 'json')
            .end((err, res) => {
                console.log(res.body)
                this.setState({
                    poos: res.body.data
                })
            });
    }

    flush(e) {
        console.log(e)
    }

    render() {

        let poos = this.state.poos.map((p) => {
            return (
                <div key={p.ID}>
                    <span className="created-at">{p.CreatedAt}</span>
                    <span className="content">{p.Content}</span>
                </div>
            )
        })


        return (
            <div>
                <div>
                    <textarea placeholder="don't even bother typing something in here, no one cares...."></textarea>
                    <button className="btn btn-primary" onClick={this.flush} id="submit">Flush</button>
                </div>
                {poos}
            </div>
        );
    }
}
render(<App />, document.getElementById("app"));