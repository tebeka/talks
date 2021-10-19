import pandas as pd

df = pd.read_csv('titanic.csv')

feature_cols = ['Pclass', 'Sex', 'Age', 'SibSp', 'Parch', 'Fare', 'Embarked']
dfX = df[feature_cols].copy()

dfX['Age'].fillna(dfX['Age'].mean(), inplace=True)

embarked = dfX.pop('Embarked')
embarked.fillna(embarked.mode()[0], inplace=True)

embarked_dummies = pd.get_dummies(embarked)

gender = dfX.pop('Sex')
gender_dummies = pd.get_dummies(gender)

dfX = pd.concat([dfX, embarked_dummies, gender_dummies], axis=1)

X = dfX.values
y = df['Survived'].values
