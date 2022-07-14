import { BrowserRouter, Routes, Route } from "react-router-dom";
import Modal from 'react-modal'
import './App.css';

// page
import RegularScheduleConfig from './components/page/regular_schedule_config';
import Authentication from './components/page/authentication';
import GroupsIndex from './components/page/groups';
import GroupPage from './components/page/group';

// context
import { GroupsContextProvider } from './context/groups'

Modal.setAppElement('#root')

function App() {

  return (
    <BrowserRouter>
        <Routes>
          <Route path="/regular_schedule_config" element={
            <Authentication>
              <RegularScheduleConfig />
            </Authentication>
          } />
          <Route path="/groups" element={
            <Authentication>
              <GroupsContextProvider>
                <GroupsIndex/>
              </GroupsContextProvider>
            </Authentication>
          } />
          <Route path="/groups/:id" element={
            <Authentication>
              <GroupsContextProvider>
                <GroupPage/>
              </GroupsContextProvider>
            </Authentication>
          }/>
        </Routes>
    </BrowserRouter>
  );
}

export default App;
