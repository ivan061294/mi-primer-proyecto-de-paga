import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import LoadingSpinner from "./LoadingSpinner";


class Content extends Component {
    componentDidMount() {
        const { browserTitle } = this.props;
        if (browserTitle) { document.title = browserTitle; }
    }

    render() {
        const {
            title, subTitle, homeRoute = '/', children, breadCrumb, loaded
        } = this.props;
        return (
            <React.Fragment>
                <section className="content-header">
                    <h1>
                        {title}
                        {' '}
                        {subTitle ? <small>{subTitle}</small> : ''}
                        {' '}
                        {!loaded?<LoadingSpinner />:null}
                    </h1>
                    <ol className="breadcrumb">
                        <li>
                            <Link to={homeRoute}>
                                <FontAwesomeIcon icon={['fas', 'tachometer-alt']} />
                                {' Home'}
                            </Link>
                        </li>
                        {breadCrumb?breadCrumb.map((bread, idx) => {
                            if (bread.link==null) {
                                return(<li key={idx} className="active">{bread.title}</li>)
                            } else {
                                return (<li key={idx}>
                                    <Link to={bread.link}>
                                        {bread.title}
                                    </Link>
                                </li>)
                            }
                        }):(<li className="active">{title}</li>)}
                    </ol>
                </section>
                <section className="content">
                    {loaded?children:null}
                </section>
            </React.Fragment>
        );
    }
}

export default Content;
