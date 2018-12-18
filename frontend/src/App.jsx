import React, { Fragment } from 'react';
import { Container, Row, Col } from 'reactstrap';
import {BrowserRouter as Router, Route, Link} from 'react-router-dom';

import Vaccine from './components/Vaccine/Vaccine';
import UserVaccines from './components/User/MyVaccine';
import UsersList from './components/User/UserList';
import Home from './components/Layout/Home';
import AboutUs from './components/Layout/AboutUs';
import Header from './components/Layout/Header';
import SideCard from './components/Layout/SideCard';
import Footer from './components/Layout/Footer';

const isLoggedIn = true;

const App = () => (
  <Fragment>
    <Header />
    <Router>
    
    <main className="my-5 py-5">
      <Route exact path="/" component={HomePage}/>
      <Route path="/vaccines" component={VaccinePage}/>
      <Route path="/about-us" component={AboutUsPage}/>
      <Route path="/user/my-vaccines" component={UserVacinnesPage}/>
      <Route path="/users" component={UsersPage}/>
    </main>

    </Router>
    <Footer />
  </Fragment>
);

function HomePage()
 {
    if (isLoggedIn)
    {
      return (
        <Container className="px-0">
      <Row noGutters className="pt-2 pt-md-5 w-100 px-4 px-xl-0 position-relative">
          <Col xs={{ order: 2 }} md={{ size: 4, order: 1 }} tag="aside" className="pb-5 mb-5 pb-md-0 mb-md-0 mx-auto mx-md-0">
          <SideCard />
        </Col>
        <Col xs={{ order: 1 }} md={{ size: 7, offset: 1 }} tag="section" className="py-5 mb-5 py-md-0 mb-md-0">
          <Home />
        </Col>
        </Row>
        </Container>
      );
    } else {
      return(
        <Container className="px-0">
      <Home />
        </Container>
       );
    }
 }

 function VaccinePage()
 {
  if (isLoggedIn)
  {
    return (
      <Container className="px-0">
    <Row noGutters className="pt-2 pt-md-5 w-100 px-4 px-xl-0 position-relative">
        <Col xs={{ order: 2 }} md={{ size: 4, order: 1 }} tag="aside" className="pb-5 mb-5 pb-md-0 mb-md-0 mx-auto mx-md-0">
        <SideCard />
      </Col>
      <Col xs={{ order: 1 }} md={{ size: 7, offset: 1 }} tag="section" className="py-5 mb-5 py-md-0 mb-md-0">
        <Vaccine />
      </Col>
      </Row>
      </Container>
    );
  } else {
    return(
      <Container className="px-0">
    <Vaccine />
      </Container>
     );
  }
 }

 function AboutUsPage()
 {
  if (isLoggedIn)
  {
    return (
      <Container className="px-0">
    <Row noGutters className="pt-2 pt-md-5 w-100 px-4 px-xl-0 position-relative">
        <Col xs={{ order: 2 }} md={{ size: 4, order: 1 }} tag="aside" className="pb-5 mb-5 pb-md-0 mb-md-0 mx-auto mx-md-0">
        <SideCard />
      </Col>
      <Col xs={{ order: 1 }} md={{ size: 7, offset: 1 }} tag="section" className="py-5 mb-5 py-md-0 mb-md-0">
        <AboutUs />
      </Col>
      </Row>
      </Container>
    );
  } else {
    return(
      <Container className="px-0">
    <AboutUs />
      </Container>
     );
  }
 }

 function UserVacinnesPage()
 {
   return(
    <Container className="px-0">
    <Row noGutters className="pt-2 pt-md-5 w-100 px-4 px-xl-0 position-relative">
    <Col xs={{ order: 2 }} md={{ size: 4, order: 1 }} tag="aside" className="pb-5 mb-5 pb-md-0 mb-md-0 mx-auto mx-md-0">
      <SideCard />
    </Col>
    <Col xs={{ order: 1 }} md={{ size: 7, offset: 1 }} tag="section" className="py-5 mb-5 py-md-0 mb-md-0">
      <UserVaccines />
    </Col>
    </Row>
    </Container>
   );
 }

 function UsersPage() {
   return (
    <Container className="px-0">
    <Row noGutters className="pt-2 pt-md-5 w-100 px-4 px-xl-0 position-relative">
    <Col xs={{ order: 2 }} md={{ size: 4, order: 1 }} tag="aside" className="pb-5 mb-5 pb-md-0 mb-md-0 mx-auto mx-md-0">
      <SideCard />
    </Col>
    <Col xs={{ order: 1 }} md={{ size: 7, offset: 1 }} tag="section" className="py-5 mb-5 py-md-0 mb-md-0">
      <UsersList />
    </Col>
    </Row>
    </Container>
   );
 }

export default App;
