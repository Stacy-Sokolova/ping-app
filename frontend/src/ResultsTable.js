import React from 'react';
import { Table, Header, Segment, Label } from 'semantic-ui-react'

export default function ResultsTable({results}) {
    const rows = results.map(((result, index) => {
        let color;

        if (index === 0) {
            color='yellow';
        } else if (index === 1) {
            color='grey';
        } else if (index === 2) {
            color='orange';
        }

        return (
            <Table.Row key={ index }>
                <Table.Cell>{ result.id }</Table.Cell>
                <Table.Cell>{ result.port }</Table.Cell>
                <Table.Cell>{ result.timeOfPing }</Table.Cell>
                <Table.Cell>{ result.dateOfPing }</Table.Cell>
            </Table.Row>
        );
    }));

    return (
        <div className="ui container">
            <Segment>
                <Table striped>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Id</Table.HeaderCell>
                            <Table.HeaderCell>Address</Table.HeaderCell>
                            <Table.HeaderCell>Time</Table.HeaderCell>
                            <Table.HeaderCell>Date</Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>
                        { rows }
                    </Table.Body>
                </Table>
            </Segment>
        </div>
    );
}