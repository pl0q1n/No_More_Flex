import React, {Component} from 'react'
import {NavLink} from 'react-router-dom'
const links = [
    {to: '/', label: 'Home page', exact: 'true' },
    {to: '/auth', label: 'Autorization', exact: 'false' },
    {to: '/add', label: 'Add transaction', exact: 'false' },
    {to: '/get', label: 'Show transactions', exact: 'false' }

];

class LinkBar extends Component {

    renderLinks() {
        return links.map((link, index) => {
            return(
                <li key={index}>
                    <NavLink
                        to={link.to}
                        exact={link.exact}
                    >
                    {link.label}
                    </NavLink>
                </li>
            )
        })
    }
    render() {
        return(
            <nav>
                <ul>
                    { this.renderLinks() }
                </ul>
            </nav>
        )
    }

}

export default LinkBar