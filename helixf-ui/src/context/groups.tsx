import React, { useState, createContext } from 'react'

// model
import { Group } from '../components/model/group'

type GroupsContextType = {
  groups: Group[]
  setGroups: React.Dispatch<React.SetStateAction<Group[]>>
}

export const GroupsContext = createContext({} as GroupsContextType)

export const GroupsContextProvider: React.FC<{ children?: JSX.Element }> = ({ children }) => {
  const [groups, setGroups] = useState([] as Group[])
  return (
    <>
      <GroupsContext.Provider value={{ groups, setGroups }}>{children}</GroupsContext.Provider>
    </>
  )
}
