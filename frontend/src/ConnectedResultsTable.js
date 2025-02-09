import React from 'react';
import ResultsTable from './ResultsTable';
import Pusher from 'pusher-js';

const socket = new Pusher(process.env.REACT_APP_PUSHER_KEY, {
    cluster: process.env.REACT_APP_PUSHER_CLUSTER,
    encrypted: true
});

export default class ConnectedResultsTable extends React.Component {
    state = {
        results: []
    };

    componentDidMount() {
        fetch('http://127.0.0.1:8080/', {
            method: 'GET'
        })
        .then((response) => response.json())
        .then((response) => this.setState(response));

        const channel = socket.subscribe('results');
        channel.bind('results', (data) => {
            this.setState(data);
        });
    }
    render() {
        return <ResultsTable results={this.state.results} />;
    }
}