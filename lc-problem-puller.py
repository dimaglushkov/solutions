import argparse
import leetcode
import os
import pandas as pd

lang_specifics = {
    'golang': {
        'ext': 'go',
        'com': '//',
        'prefix': 'package main\n\n'
    },
    'python': {
        'ext': 'py',
        'com': '#'
    }
}


def get_title_slugs(problems: list, lang: str, d: str) -> dict:
    slugs = dict()
    for problem in problems:
        if problem.startswith('http'):
            slug = problem
            if problem.endswith('/'):
                slug = slug[:-1].split('/')[-1]
            slugs[slug] = os.path.join(d, f'{slug}.{lang_specifics[lang]["ext"]}')
        else:
            slugs[problem] = os.path.join(d, f'{problem}.{lang_specifics[lang]["ext"]}')

    existing_solutions = list()
    for k, v in slugs.items():
        if os.path.exists(v):
            print(f'Skipping {k}, {v} already exists. ')
            existing_solutions.append(k)
    for i in existing_solutions:
        del slugs[i]

    return slugs


def get_slug_data(name: str) -> dict:
    configuration = leetcode.Configuration()
    configuration.api_key["Referer"] = "https://leetcode.com"
    configuration.debug = False
    api_instance = leetcode.DefaultApi(leetcode.ApiClient(configuration))
    graphql_request = leetcode.GraphqlQuery(
        query="""
                query getQuestionDetail($titleSlug: String!) {
                  question(titleSlug: $titleSlug) {
                    questionId
                    questionFrontendId
                    boundTopicId
                    title
                    content
                    translatedTitle
                    difficulty
                    similarQuestions
                    topicTags {
                      name
                      slug
                    }
                    codeSnippets {
                      lang
                      langSlug
                      code
                      __typename
                    }
                    stats
                    codeDefinition
                    hints
                    solution {
                      id
                      canSeeDetail
                      __typename
                    }
                    status
                    sampleTestCase
                    enableRunCode
                    metaData
                    translatedContent
                    judgerAvailable
                    judgeType
                    mysqlSchemas
                    enableTestMode
                    envInfo
                    __typename
                  }
                }
            """,
        variables=leetcode.GraphqlQueryGetQuestionDetailVariables(title_slug=name),
        operation_name="getQuestionDetail",
    )
    return api_instance.graphql_post(body=graphql_request).to_dict()


def create_code_template(slug: str, file: str, lang: str, data: dict) -> None:
    source_link = f'https://leetcode.com/problems/{slug}/'
    code_snippet = ''
    for snippet in data['code_snippets']:
        if snippet['lang_slug'] == lang:
            code_snippet = snippet['code']
            break

    code = ''
    if 'prefix' in lang_specifics[lang].keys():
        code = lang_specifics[lang]['prefix']

    code += f'{lang_specifics[lang]["com"]} source: {source_link}\n\n'
    code += code_snippet
    with open(file, 'w') as code_file:
        code_file.write(code)


def update_meta_file(slug: str, lang: str, slug_data: dict) -> None:
    data = pd.read_csv('leetcode/.meta.csv', index_col='slug').to_dict('index')

    solutions = {f.split('.')[0]: f.split('.')[1] for f in os.listdir('leetcode') if not f.startswith('.')}

    if slug not in data.keys() or data['slug']['lang'] != solutions[slug]:
        data[slug] = {
            'id': slug_data['question_frontend_id'],
            'difficulty': slug_data['difficulty'],
            'tags': [tag['name'] for tag in slug_data['topic_tags']],
            'lang': lang
        }
        pd.DataFrame.from_dict(data).transpose().to_csv('leetcode/.meta.csv', index_label='slug', index=True)

    for problem in data.keys():
        # keeping meta.csv actual by removing problems with no solution, but it's not much likeable to happen
        if problem not in solutions:
            del data[problem]
            pd.DataFrame.from_dict(data).transpose().to_csv('leetcode/.meta.csv', index_label='slug', index=True)




def main():
    parser = argparse.ArgumentParser(description="Pulls problems from leetcode and sets up env in local repo")
    parser.add_argument('problems', type=str, nargs='+', help='Name or link to the problems')
    parser.add_argument('--lang', dest='lang', action='store', default='golang', help='Solution lang')
    parser.add_argument('--dir', dest='dir', action='store', default='leetcode', help='Solutions directory')
    args = parser.parse_args()

    slugs = get_title_slugs(args.problems, args.lang, args.dir)
    if len(slugs) == 0:
        print("Nothing to do")
        return

    for slug, file in slugs.items():
        slug_data = get_slug_data(slug)['data']['question']
        create_code_template(slug, file, args.lang, slug_data)
        update_meta_file(slug, args.lang, slug_data)


if __name__ == '__main__':
    main()
