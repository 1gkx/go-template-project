# Порядок внесения изменений в проект

#### Содержание:
- [Code style](/CONTRIBUTING.md#code-style)
- [Code review](/CONTRIBUTING.md#code-review)
- [Безопасность](/CONTRIBUTING.md#%D0%B1%D0%B5%D0%B7%D0%BE%D0%BF%D0%B0%D1%81%D0%BD%D0%BE%D1%81%D1%82%D1%8C)
- [Общие правила](/CONTRIBUTING.md#%D0%BE%D0%B1%D1%89%D0%B8%D0%B5-%D0%BF%D1%80%D0%B0%D0%B2%D0%B8%D0%BB%D0%B0)

## Оформление коммитов
Рекомендуется следовать [общим рекомендациям][Common git messages src] к сообщениям git-коммитов и оформлять их в
соответствии с правилами [Conventional Commits][Conventional commits src]. В футере коммита (если это не hotfix) должен
содержаться номер задачи в Jira:
```
docs: Add CONTRIBUTING.md

Optional detailed description.

MKB11-55
```

Если не требуется CI проверок для merge-request, можно в футере коммита указать:
```
chore: Update version

[ci-skip]
```

## Code style
В проекте рекомендуется использовать [Google Go Style][Google Go Style src].

## Code review
- При реализации новых API-методов следует использовать REST-подход. Рекомендуется использовать практики описанные в [Google API design guide][Google API design guide].
- Проводя ревью, следует руководствоваться [официальными рекомендациями к Go-ревью][Review comments].

[Conventional commits src]: https://conventionalcommits.org
[Common git messages src]: https://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
[Review comments]: https://github.com/golang/go/wiki/CodeReviewComments
[Google Go Style src]: https://google.github.io/styleguide/go
[Google API design guide]: https://cloud.google.com/apis/design