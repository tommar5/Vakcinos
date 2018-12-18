import React, { Fragment } from 'react';

import {
  Button, Card, CardImg, CardBody,
  CardTitle, CardSubtitle, CardText
} from 'reactstrap';

import {Link} from 'react-router-dom';

import EditProfile from '../User/EditProfileModal';

const BANNER = 'https://i.imgur.com/CaKdFMq.jpg';

const SideCard = () => (
  <Fragment>
    
    <Card>
      <CardImg top width="100%" src={BANNER} alt="banner" />
      <CardBody>
        <CardTitle className="h3 mb-2 pt-2 font-weight-bold text-secondary">Glad Chinda</CardTitle>
        <CardSubtitle className="text-secondary mb-3 font-weight-light text-uppercase" style={{ fontSize: '0.8rem' }}>Web Developer, Lagos</CardSubtitle>
        <CardText className="text-secondary mb-4" style={{ fontSize: '0.75rem' }}>Full-stack web developer learning new hacks one day at a time. Web technology enthusiast. Hacking stuffs @theflutterwave.</CardText>
        <Link to="/users"><Button color="info" outline className="font-weight-bold">Users</Button></Link>{' '}
        <Link to="/user/my-vaccines"><Button color="info" outline className="font-weight-bold">My vaccines</Button></Link>{' '}
        <EditProfile/>
        <Link to="/logout"><Button color="secondary" outline className="font-weight-bold">Logout</Button></Link>
      </CardBody>
    </Card>
    
  </Fragment>
);

export default SideCard;