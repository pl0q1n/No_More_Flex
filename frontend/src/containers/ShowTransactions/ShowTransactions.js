import React, { Component } from 'react'


import TransactionPie from '../../components/Diagrams/TransactionPie/TransactionPie'

class ShowTransactions extends Component {
	constructor(props) {
		super(props)
		this.config = {
	    	headers: {
	    		'Accept': 'application/json',
    			'Content-Type': 'application/json',
	    	}
		};
	}
	
	render() {
			return (
				<div>
					<TransactionPie/>
				</div>
			)
		}
}

export default ShowTransactions