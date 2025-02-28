# Contributing Guidelines

To help everyone get involved in the project, we have outlined the following contribution guidelines.

## Getting Started
1.	Fork the repository.
2.	Clone your forked repository to your local machine.
3.	Create a new branch. Typically, branches should start with `feat/` for new features or `fix/` for bug fixes.
4.	Develop your changes in the new branch.
5.	Before submitting, pull the latest changes from the original repository to ensure no conflicts.
6.	Push your changes to your forked repository.
7.	Create a Pull Request to the `main` branch of the original repository.

## Commit Guidelines
- Code Style: Please follow the project’s coding style. If you are unsure, refer to the existing code.
- Unit Tests: Ensure that your code passes all unit tests.
- Commit Message: Please write commit messages in English, following the format `type(scope): description`, for example, `feat(client): add new feature`, `fix(server): fix a bug`, `docs: update README.md`, etc.

## Branches and Commit Messages

Both branches and commit messages should include a prefix, which helps quickly identify the type and scope of the change.

For this, we define the following conventions:

Types:
- feat: New features
- fix: Bug fixes
- docs: Documentation-related changes
- style: Code style-related changes
- refactor: Code refactoring
- test: Testing-related changes
- chore: Miscellaneous changes
- revert: Reverts a previous commit

Scopes:
- client: Refers to changes related to the client side
- server: Refers to changes related to the server side

For development-related commits, the scope is usually required, e.g., `feat(client): add new feature`, whereas for documentation or miscellaneous changes, the scope is often unnecessary, e.g., `docs: update README.md`.
