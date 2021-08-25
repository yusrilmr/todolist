# Todolist Backend
The backend of todolist app.

## Design Decision
- One Task can have multiple Labels
- One Label can have multiple Tasks
- 

## Compromise
- By default, GORM implements Upsert instead for updating the association
  - Cannot update Task's label yet. It only supports upsert.
  - Cannot update Label's task yet. It only supports upsert.