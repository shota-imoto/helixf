import { BrowserRouter, Routes, Route } from "react-router-dom";
import './App.css';

// page
import RegularScheduleTemplateConfig from './components/page/config';
import Authentication from './components/page/authentication';
import GroupsIndex from './components/page/groups';
import GroupPage from './components/page/group';

// context
import { GroupsContextProvider } from './context/groups'



function App() {

  return (
    <BrowserRouter>
        <Routes>
          <Route path="/config" element={
            <Authentication>
              <RegularScheduleTemplateConfig />
            </Authentication>
          } />
          <Route path="/groups" element={
            <Authentication>
              <GroupsContextProvider>
                <GroupsIndex/>
              </GroupsContextProvider>
            </Authentication>
          } />
          <Route path="/group/:id" element={
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
