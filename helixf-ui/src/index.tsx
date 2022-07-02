import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import reportWebVitals from './reportWebVitals';
import RegularScheduleTemplateConfig from './components/page/config';
import Authentication from './components/page/authentication';
import GroupsIndex from './components/page/groups';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.Fragment>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
          <Route path="/config" element={
            <Authentication>
              <RegularScheduleTemplateConfig />
            </Authentication>
          } />
          <Route path="/groups" element={
            <Authentication>
              <GroupsIndex/>
            </Authentication>
          } />
      </Routes>
    </BrowserRouter>
  </React.Fragment>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
