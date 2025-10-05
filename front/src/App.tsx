import React from 'react';
import OwnerDashboard from './pages/OwnerDashboard';
import Header from './layouts/Header';
import GlobalNav from './layouts/GlobalNav';

function App() {
  return (
    <div className="App">
      <Header />
      <GlobalNav />
      <OwnerDashboard />
    </div>
  )
}

export default App;
