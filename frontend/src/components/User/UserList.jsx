import React, { Component, Fragment } from 'react';
import { Link } from 'react-router-dom';
import { Row, Col, ListGroup, ListGroupItem, ListGroupItemHeading, ListGroupItemText, UncontrolledCollapse, Button, CardBody, Card } from 'reactstrap';

import UpdateUser from './UpdateUser';

class UserList extends Component {

  state = { vaccine: null }
  
  constructor(props){
    super(props);
    this.state = {
      vaccines: [],
      isLoaded: false,

    }
  }

  componentDidMount() {
    fetch('https://jsonplaceholder.typicode.com/users')
    .then(res => res.json())
    .then(json => {
      this.setState({
        isLoaded: true,
        vaccine: json,
      })
    });
  }
  
  render() {

    var { isLoaded, vaccines } = this.state;

    if(!isLoaded) {
      return <div>Loading...</div>;
    }
    else {
      return (
        <Fragment>
        <div className="position-relative">
 
            <span className="pb-4 h2 text-dark border-bottom border-gray">All users</span>

          <ListGroup>
        <ListGroupItem>
          <a href="" id="toggler" style={{ marginBottom: '1rem' }}>
            <ListGroupItemHeading>
              Name SecondName
            </ListGroupItemHeading>
          </a>
        
          <ListGroupItemText>
              <Row>
                  <Col>
                  BirthDay: ... Email: test@test.com
                  </Col>
                  <Col>
                    <div className="align-right">
                        <UpdateUser/>
                        <Link to="user/{id}/delete"><Button color="danger" style={{marginRight: 15}}>Delete</Button></Link>
                    </div>
                  </Col>
              </Row>
          </ListGroupItemText>
          <UncontrolledCollapse toggler="#toggler">
      <Card>
        <CardBody>
          <Row>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Nesciunt magni, voluptas debitis
          similique porro a molestias consequuntur earum odio officiis natus, amet hic, iste sed
          dignissimos esse fuga! Minus, alias.
          </Row>
        </CardBody>
      </Card>
    </UncontrolledCollapse>
        </ListGroupItem>
      </ListGroup>
          
        </div>
      </Fragment>
      );
    }
  }
  
}

export default UserList;