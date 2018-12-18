import React from 'react';
import { Jumbotron, Button } from 'reactstrap';
import {Link} from 'react-router-dom';

const Home = () => {
    return (
        <div>
            <Jumbotron>
                <h1 className="display-3">Vaccinate</h1>
                <p className="lead">Vaccination is the most effective prophylaxis from tick-borne encephalitis.</p>
                <hr className="my-2" />
                <p>We vaccinate all vaccines registered in the European Union. For more information contact with us or check available vaccines.</p>
                <p className="lead">
                <Link to="/vaccines"><Button color="info">Check vaccines</Button></Link>
                </p>
            </Jumbotron>
        </div>
    );
}

export default Home;